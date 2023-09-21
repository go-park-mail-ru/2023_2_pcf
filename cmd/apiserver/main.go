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
	flag.Parse() // Парсим флаг с путем до конфига

	config := apiserver.NewConfig()               // Создаем новый конфиг с дефолтными значениями
	_, err := toml.DecodeFile(configPath, config) // Парсим файл в наш созданный конфиг
	if err != nil {
		// Здесь будет лог с уровнем еррор и отлов паники
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		// Здесь будет лог с уровнем еррор и отлов паники
	}
}
