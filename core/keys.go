package core

const (
	keyLimit = 6
)

type Key interface {
	Token() string
	Extend(string)
	Metadata() interface{}
}

type AWSKey struct {
	bucket string
	key    string
}

func (a *AWSKey) Token() string {
	return a.key
}

func (a *AWSKey) Extend(subKey string) {
	a.key += "/" + subKey
}

func (a *AWSKey) Metadata() interface{} {
	return a.bucket
}

func NewS3Key(bucket string, key string) *AWSKey {
	return &AWSKey{bucket, key}
}
