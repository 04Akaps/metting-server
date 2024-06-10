package service

import (
	util "github.com/04Akaps/go-util/log"
	"github.com/04Akaps/metting/aws"
	"github.com/04Akaps/metting/config"
	"github.com/04Akaps/metting/db/HDD"
	. "github.com/04Akaps/metting/db/HDD/types"
	"github.com/04Akaps/metting/db/SSD"
	"github.com/04Akaps/metting/module/API/types"
	"mime/multipart"
	"net/http"
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
	LikeSomeOne(from, to string) error
	RefuseLike(from, to string) error
	AcceptedLike(from, to string) error
	GetLikedList(user string) ([]*User, error)
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
