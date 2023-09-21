package main

import (
	"AdHub/internal/app/apiserver"
	"flag"

	"github.com/BurntSushi/toml"
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

	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {

	}

}
