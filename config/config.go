package config

import "github.com/eric-lindau/flip/core"

type Env struct {
	MaxData   int
	DataStore core.DataStore
	KeyFunc   func(string, *core.KeyOptions) core.Key
}
