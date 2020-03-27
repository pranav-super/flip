package config

import "github.com/eric-lindau/flip/core"

type Env struct {
	MaxData int
	Flip    core.Flip
	KeyFunc func(string, *core.KeyOptions) core.Key
}
