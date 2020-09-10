package main

import (
	"flag"
	"github.com/nikionov/job-backend-trainee/internal/app/apiserver"
	"github.com/nikionov/job-backend-trainee/internal/app/apiserver/pkg/mod/github.com/burntsushi/toml@v0.3.1"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}