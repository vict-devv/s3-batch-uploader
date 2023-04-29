package aws

import (
	"mime"
	"os"
	"path/filepath"
	"strings"
	"vict-devv/s3-batch-uploader/constants"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type FileInfo struct {
	key      string
	fullPath string
}

type SyncFolderIterator struct {
	bucket    string
	fileInfos []FileInfo
	err       error
}

func (s *SyncFolderIterator) Next() bool {
	return len(s.fileInfos) > 0
}

func (s *SyncFolderIterator) Err() error {
	return s.err
}

func (s *SyncFolderIterator) UploadObject() s3manager.BatchUploadObject {
	fi := s.fileInfos[0]
	s.fileInfos = s.fileInfos[1:]

	body, err := os.Open(fi.fullPath)
	if err != nil {
		s.err = err
	}

	extension := filepath.Ext(fi.key)
	mimeType := mime.TypeByExtension(extension)
	if mimeType == "" {
		mimeType = constants.DefaultMimeType
	}

	input := s3manager.UploadInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(fi.key),
		Body:        body,
		ContentType: aws.String(mimeType),
	}

	return s3manager.BatchUploadObject{
		Object: &input,
	}
}

func NewSyncFolderIterator(path, bucket string) (*SyncFolderIterator, error) {
	var metadata []FileInfo

	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			key := strings.TrimPrefix(p, path)
			metadata = append(metadata, FileInfo{key: key, fullPath: p})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &SyncFolderIterator{
		bucket,
		metadata,
		nil,
	}, nil
}
