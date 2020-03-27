package core

import (
	"io"
)

// Generic Flip interface.
// Variants might use multiple Storage mechanisms/strategies.
// Multiple objects are stored by name under a single Key. // TODO: Impl according to this
type Flip interface {
	Objects(key string) ([]string, error)
    //GenerateKey(opt *KeyOptions) (string, error)
	Get(key string, obj string) ([]byte, error)
	Put(key string, obj string, r io.Reader) error
}

// Manages actual key retrieval and storage mechanisms.
// Abstraction allows for e.g. cascade from memory -> S3.
type Storage interface {
    SubObjects(location string) ([]string, error)
    Write(key string, data io.Reader) error
    Read(key string) ([]byte, error)
}

// Flip impl. utilizing only a single storage mechanism.
type FlipStore struct {
    storage Storage
}

func (flip *FlipStore) Objects(key string) ([]string, error) {
    return flip.storage.SubObjects(key)
}

func (flip *FlipStore) Put(data io.Reader, dst string) error {
	return flip.storage.Write(dst, data)
}

func (flip *FlipStore) Get(key string) ([]byte, error) {
    return flip.storage.Read(key)
}

// Generates Key for client data retrieval
//func GenerateKey(keyFunc func(string, *KeyOptions) Key, options *KeyOptions) (Key, error) {
//	seed := time.Now().UnixNano()
//	source := rand.NewSource(seed)
//	r := rand.New(source)
//	id := strconv.Itoa(r.Intn(int(math.Pow10(keyLimit))))
//
//	return keyFunc(id, options), nil // TODO: Return/check error?
//}
