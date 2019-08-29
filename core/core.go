package core

import (
	"io"
	"time"
	"math/rand"
	"strconv"
	"math"
)

type KeyOptions struct {
	TTL int
}

// TODO: Figure out empty name
func PutData(store DataStore, data io.Reader, name string, dst Key) error {
	store.putData(dst.Extend(name), data) // TODO: Impl./return error
	return nil
}

func GetData(store DataStore, key Key) []byte {
	data, _ := store.getData(key)
	return data // TODO: Impl./return error
}

// Generates Key for client data retrieval
func GenerateKey(keyFunc func(string, *KeyOptions) Key, options *KeyOptions) (Key, error) {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	r := rand.New(source)
	id := strconv.Itoa(r.Intn(int(math.Pow10(keyLimit))))

	return keyFunc(id, options), nil // TODO: Return/check error?
}
