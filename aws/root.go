package aws

import (
	"github.com/04Akaps/go-util/log"
	"github.com/04Akaps/metting/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type GoAWS struct {
	session    *session.Session
	S3         *s3.S3
	S3Uploader *s3manager.Uploader
	Bucket     string

	cfg *config.Config
	log *log.Log
}

func NewAWS(cfg *config.Config, log *log.Log) *GoAWS {
	g := &GoAWS{cfg: cfg, log: log}

	var err error

	awsCfg := cfg.Aws

	if g.session, err = session.NewSession(&aws.Config{
		Region:      aws.String(awsCfg.Region),
		Credentials: credentials.NewStaticCredentials(awsCfg.IAMAKey, awsCfg.IAMSKey, ""),
	}); err != nil {
		panic(err)
	} else {
		g.Bucket = awsCfg.Bucket
		g.S3 = s3.New(g.session)
		g.S3Uploader = s3manager.NewUploader(g.session)

		return g
	}
}
