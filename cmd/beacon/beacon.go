package main

import (
	"github.com/eric-lindau/flip/config"
	"github.com/eric-lindau/flip/http"
)

func main() {
	env := &config.Env{MaxData: 500}
	http.Init(env)
}
