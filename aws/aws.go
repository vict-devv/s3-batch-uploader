package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func NewAWSSession(region, profile string) *session.Session {
	return session.Must(session.NewSessionWithOptions(
		session.Options{
			Profile: profile,
			Config: aws.Config{
				Region: aws.String(region),
			},
		},
	))
}

func UploadFolderToS3(sess *session.Session, path, bucket string) error {
	iter, err := NewSyncFolderIterator(path, bucket)
	if err != nil {
		return err
	}

	uploader := s3manager.NewUploader(sess)

	err = uploader.UploadWithIterator(aws.BackgroundContext(), iter)
	if err != nil {
		return err
	}

	if iter.err != nil {
		return iter.err
	}

	return nil
}
