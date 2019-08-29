package core

const (
	keyLimit = 6
)

type Key interface {
	Token() string
	Extend(string) Key
	Metadata() interface{}
}

type AWSKey struct {
	bucket string
	key    string
}

func (a *AWSKey) Token() string {
	return a.key
}

func (a *AWSKey) Extend(subKey string) Key {
	return &AWSKey{
		bucket: a.bucket,
		key: a.key + "/" + subKey,
	}
}

func (a *AWSKey) Metadata() interface{} {
	return a.bucket
}

// TODO: Implement Keys in diff. buckets - map key/token to bucket/key (Dynamo)
func NewS3Key(bucket string, key string) *AWSKey {
	return &AWSKey{"s3.flip.io", key}
}
