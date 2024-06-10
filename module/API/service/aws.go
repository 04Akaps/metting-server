package service

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"os"
)

func (s *service) putFileToS3(fileName, userName, withoutDot, filePath string) error {
	fileKey := userName + "/" + fileName

	if f, err := os.Open(filePath); err != nil {
		return err
	} else {
		defer f.Close()

		if fileState, err := f.Stat(); err != nil {
			return err
		} else {
			if fileState.Size() <= 100000000 {
				if err = s.putFileToS3UsingPutObject(fileKey, withoutDot, f); err != nil {
					return err
				}
			} else {
				if err = s.pubFileToS3UsingUploader(fileKey, withoutDot, f); err != nil {
					return err
				}
			}
		}

		return nil
	}
}

func (s *service) putFileToS3UsingPutObject(fileKey, withoutDot string, file *os.File) error {

	//특정 객체를 업로드 하는 함수
	uploadInput := &s3.PutObjectInput{
		Bucket:      aws.String(s.aws.Bucket),
		Key:         aws.String(fileKey),
		Body:        file,
		ContentType: aws.String(fmt.Sprintf("%s/%s", "image", withoutDot)),
		ACL:         aws.String("public-read"),
	}

	_, err := s.aws.S3.PutObject(uploadInput)

	return err
}

func (s *service) pubFileToS3UsingUploader(fileKey, withoutDot string, file *os.File) error {
	// Uploader와 PutObject의 차이는 PutObject는 작은 용량의 파일을 업로드 할떄 유리
	// Uploader내부적으로 파일을 쪼개서 업로드 하기 떄문에 큰 용량의 파일을 업로드 할 떄 유리하다.
	// 대략적으로 100MB을 기준으로 사용
	uploadInput := &s3manager.UploadInput{
		Bucket:      aws.String(s.aws.Bucket),
		Key:         aws.String(fileKey),
		Body:        file,
		ContentType: aws.String(fmt.Sprintf("%s/%s", "image", withoutDot)),
		ACL:         aws.String("public-read"),
	}

	_, err := s.aws.S3Uploader.Upload(uploadInput)

	return err
}

func (s *service) GetFileFromS3(bucket, key string) ([]byte, error) {
	downloadInput := &s3.GetObjectInput{
		//Bucket: aws.String(bucket),
		Key: aws.String(key),
	}

	if res, err := s.aws.S3.GetObject(downloadInput); err != nil {
		return nil, err
	} else {
		defer res.Body.Close()

		if body, err := io.ReadAll(res.Body); err != nil {
			return nil, err
		} else {
			return body, nil
		}
	}

}

func makeMetaData(key, value string) map[string]*string {
	metadata := make(map[string]*string)

	metadata[key] = aws.String(value)
	return metadata
}
