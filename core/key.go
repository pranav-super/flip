package core

const (
	keyLimit = 6
)

type Key interface {
	Id() string
}

type AWSKey struct {
	s3Bucket string
	s3Key    string
}

func (a *AWSKey) Id() string {
	return a.s3Key
}
