package storage

import (
	"io"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	region = "us-east-1" // TODO: Decouple region
)

func keySuffix(s string) string {
	subkeys := strings.Split(s, "/")
	return subkeys[len(subkeys)-1]
}

type S3 struct {
	client     *s3service.S3
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
}

func S3Session() *S3 {
	sess := session2.Must(session2.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	return &S3{
		client:     s3service.New(sess),
		uploader:   s3manager.NewUploader(sess),
		downloader: s3manager.NewDownloader(sess),
	}
}

func (s *S3) objects(key Key) []string {
	objects, err := s.client.ListObjects(&s3service.ListObjectsInput{
		Bucket: aws.String(key.Metadata().(string)),
		Prefix: aws.String(key.Token()),
	})
	if err != nil {
		return nil // TODO: Return err
	}

	names := make([]string, len(objects.Contents))
	for i, obj := range objects.Contents {
		names[i] = keySuffix(*obj.Key)
	}

	return names
}

// TODO: Improve memory efficiency; stream to given io (e.g. echo Body) instead of into memory
func (s *S3) getData(key Key) ([]byte, bool) {
	buf := aws.NewWriteAtBuffer([]byte{})
	n, err := s.downloader.Download(buf, &s3service.GetObjectInput{
		Bucket: aws.String(key.Metadata().(string)), // Type assertion - poor performance
		Key:    aws.String(key.Token()),
	})
	if err != nil || n < 0 {
		return nil, false
	}
	return buf.Bytes(), true
}

func (s *S3) putData(key Key, r io.Reader) bool {
	_, err := s.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(key.Metadata().(string)),
		Key:    aws.String(key.Token()),
		Body:   r,
	})
	if err != nil {
		return false
	}
	return true
}
