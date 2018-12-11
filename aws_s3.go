package container_launcher

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"strings"
)

func init() {
	RegisterResolver(&s3Resolver{})
}

type s3Resolver struct {
	sess       *session.Session
	downloader *s3manager.Downloader
}

func (s s3Resolver) IsDefinedAt(str string) bool {
	return strings.HasPrefix(str, "s3://")
}

func (s s3Resolver) UsageText() string {
	return "s3://<bucketName>/<path>/<to>/<key>"
}

func (s *s3Resolver) Resolve(str string, w io.WriterAt) error {
	if s.sess == nil {
		sess, err := session.NewSession()
		if err != nil {
			return err
		}
		s.sess = sess
		s.downloader = s3manager.NewDownloader(s.sess)
	}
	trimmed := strings.TrimPrefix(str, "s3://")
	parts := strings.SplitN(trimmed, "/", 2)
	if len(parts) != 2 {
		return errors.New(fmt.Sprintf("Invalid S3 URI '%s' expected 's3://<bucket>/<key>'", str))
	}
	_, err := s.downloader.Download(w, &s3.GetObjectInput{
		Bucket: aws.String(parts[0]),
		Key:    aws.String(parts[1]),
	})
	return err
}
