package core

import (
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	session2 "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	s4 "github.com/aws/aws-sdk-go/service/s3"
	"io"
)

const (
	region = "us-east-1" // TODO: Naming convention?
)

type DataStore interface {
	getData() bool
	putData(key Key, reader io.Reader) bool
}

type S3 struct {
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
}

func S3Session() *S3 {
	sess := session2.Must(session2.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	return &S3{
		uploader:   s3manager.NewUploader(sess),
		downloader: s3manager.NewDownloader(sess),
	}
}

// TODO: Improve memory efficiency; stream to given io (e.g. echo Body) instead of into memory
func (s S3) getData(key Key) ([]byte, bool) {
	buf := aws.NewWriteAtBuffer([]byte{})
	n, err := s.downloader.Download(buf, &s4.GetObjectInput{
		Bucket: aws.String("s3.flip.io"),
		Key:    aws.String(key.Id()),
	})
	if err != nil || n < 0 {
		return nil, false
	}
	return buf.Bytes(), true
}

func (s S3) putData(key Key, r io.Reader) bool {
	_, err := s.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("s3.flip.io"),
		Key:    aws.String(key.Id()),
		Body:   r,
	})
	if err != nil {
		return false
	}
	return true
}
