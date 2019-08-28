package core

const (
	keyLimit = 6
)

type Key interface {
	Id() string
}

type AWSKey struct {
	S3Bucket string
	S3Key    string
}

func (a *AWSKey) Id() string {
	return a.S3Key
}
