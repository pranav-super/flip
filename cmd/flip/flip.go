package main

import (
	"github.com/eric-lindau/flip/config"
	"github.com/eric-lindau/flip/http"
	"github.com/eric-lindau/flip/core"
)

func main() {
	s3 := core.S3Session()
	env := &config.Env{
		MaxData:   500,
		DataStore: s3,
		KeyFunc: func(key string, options *core.KeyOptions) core.Key { // TODO: key -> token rename globally?
			// TODO: Determine bucket based on KeyOptions
			return core.NewS3Key("s3.flip.io", key)
		},
	}
	http.Init(env)
}
