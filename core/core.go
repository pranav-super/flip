package core

import (
	"io"
	"github.com/eric-lindau/flip/config"
	"time"
	"strconv"
	"math"
	"math/rand"
	"path"
)

var (
	s3 = S3Session()
)

// TODO: Figure out empty name
func ProcessData(data io.Reader, name string, dst Key) error {
	newKey := &AWSKey{"s3.flip.io", path.Join(dst.Id(), name)}
	s3.putData(newKey, data)

	//file, err := os.Create(dst)
	//if err != nil {
	//	print(err.Error())
	//	return err
	//}
	//defer file.Close()
	//
	//if _, err := io.Copy(file, data); err != nil {
	//	print(err)
	//	return err
	//}

	return nil
}

func GetData(key Key) []byte {
	data, _ := s3.getData(key)
	return data
}

func register(env *config.Env) (Key, error) {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	r := rand.New(source)
	id := strconv.Itoa(r.Intn(int(math.Pow10(keyLimit))))

	return &AWSKey{"s3.flip.io", id}, nil
}

// Generates Key for client data retrieval
func GenerateKey(env *config.Env) (Key, error) {
	key, err := register(env)
	if err != nil {
		return nil, err
	}

	return key, err
}
