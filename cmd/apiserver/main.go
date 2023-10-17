package main

import (
	"AdHub/internal/server"
	"flag"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/", "path to config file")
}

func main() {
	flag.Parse() // Парсим флаг с путем до конфига

	err := server.Parse(configPath)
	if err != nil {
		// Здесь будет лог с уровнем еррор и отлов паники
	}
	config := server.NewConfig() // Создаем новый конфиг с дефолтными значениями
	s := server.New(config)
	if err := s.Start(); err != nil {
		// Здесь будет лог с уровнем еррор и отлов паники
	}
}
