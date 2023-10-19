package main

import (
	server "AdHub/internal/app"
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

	s := server.New(server.NewConfig())
	if err := s.Start(); err != nil {
		// Здесь будет лог с уровнем еррор и отлов паники
	}
}
