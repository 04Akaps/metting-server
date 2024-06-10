package service

import (
	"fmt"
	util "github.com/04Akaps/go-util/log"
	"github.com/04Akaps/metting/aws"
	"github.com/04Akaps/metting/config"
	"github.com/04Akaps/metting/db/HDD"
	. "github.com/04Akaps/metting/db/HDD/types"
	"github.com/04Akaps/metting/db/SSD"
	"github.com/04Akaps/metting/module/API/types"
	. "github.com/04Akaps/metting/types"
	. "github.com/04Akaps/metting/types/errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type service struct {
	cfg *config.Config
	log *util.Log

	client *http.Client

	db    *HDD.DB
	redis SSD.Redis
	aws   *aws.GoAWS
}

type ServiceImpl interface {
	RegisterUser(req types.RegisterUser) error
	UploadFile(userName string, header *multipart.FileHeader, file multipart.File) error
	FindAroundFriends(user string, searchRange, limit int64) ([]*AroundUser, error)
}

func NewService(
	cfg *config.Config,
	log *util.Log,
	db *HDD.DB,
	aws *aws.GoAWS,
	redis SSD.Redis,
) ServiceImpl {

	s := &service{
		cfg:   cfg,
		log:   log,
		db:    db,
		aws:   aws,
		redis: redis,
	}

	return s
}

func (s *service) RegisterUser(req types.RegisterUser) error {
	var retryCount = 0

createAgain:
	if err := s.db.RegisterUser(req.UserName, req.Description, req.Hobby, req.Latitude, req.Hardness); err != nil {
		// 세션을 생성해서 두개의 테이블에 생성을 하기 떄문에,
		// 혹시나 internal에서 중복된 raw에 대한 lock을 획득하지 못해
		// 데드락이 발생을 하는 상황에 대해서 방지하고자 retry 적용
		retryCount++

		if retryCount < 3 {
			goto createAgain
		} else {
			s.log.ErrLog("Failed To Create User", "user", req.UserName, "err", err.Error())
			return err
		}

	} else {
		s.log.InfoLog("Success To Create DB", "user", req.UserName)
		return nil
	}
}

func (s *service) UploadFile(userName string, header *multipart.FileHeader, file multipart.File) error {
	// 유저 체크

	if u, err := s.GetUser(userName); err != nil {
		return err
	} else if imgLink, err := s.uploadFileToS3(u.UserName, header, file); err != nil {
		return err
	} else if err = s.db.UpdateUserImage(u.UserName, imgLink); err != nil {
		return err
	} else {
		return nil
	}
}

func (s *service) FindAroundFriends(user string, searchRange, limit int64) ([]*AroundUser, error) {
	// MySQL 검색 쿼리를 사용하여 유저들의 데이터를 가져 온다.

	if limit == 0 {
		limit = 5
	}

	if u, err := s.db.GetUser(user); err != nil {
		return nil, err
	} else if users, err := s.db.GetAroundFriends(u.UserName, u.Latitude, u.Hardness, searchRange, limit); err != nil {
		return nil, err
	} else {
		return users, nil
	}

}

func (s *service) GetUser(userName string) (*User, error) {
	if u, err := s.db.GetUser(userName); err != nil {
		s.log.ErrLog("Failed To Get User", "user", userName, "err", err.Error())
		return nil, err
	} else {
		return u, nil
	}
}

func (s *service) uploadFileToS3(userName string, header *multipart.FileHeader, file multipart.File) (string, error) {
	fileName := header.Filename
	fileExt := filepath.Ext(fileName)

	if !solveImageExtension(fileExt) {
		return "", Err(FileOnlyAccept)
	} else {
		imgPath := IMG_PATH
		os.MkdirAll(imgPath, os.ModePerm)
		filePath := fmt.Sprintf("%s/%s", imgPath, fileName)

		// 파일을 저장합니다.
		if out, err := os.Create(filePath); err != nil {
			s.log.ErrLog("Failed To Create File", "err", err.Error())
			return "", err
		} else {
			defer func() {
				defer out.Close()

				if err = os.Remove(filePath); err != nil {
					// Internal에서 cron을 통해 확인을 하며 잔여 삭제 파일이 있으면 삭제 할 예정
					key := "meeting-remove-" + filePath
					if err = s.redis.Save(key, filePath, 0); err != nil {
						s.log.ErrLog("Failed To Save To Remove Data To Redis", "err", err.Error())
					}
				} else {
					s.log.InfoLog("Success Remove File", "path", filePath)
				}
			}()

			if _, err = io.Copy(out, file); err != nil {
				return "", err
			} else {
				if err = s.putFileToS3(fileName, userName, strings.TrimPrefix(fileExt, "."), filePath); err != nil {
					s.log.ErrLog("Failed To Create S3 Bucket", "err", err.Error())
					return "", err
				} else {
					imgLink := fmt.Sprintf("%s%s/%s", s.cfg.Aws.S3BaseURL, userName, fileName)

					return imgLink, nil
				}
			}
		}

	}

}
