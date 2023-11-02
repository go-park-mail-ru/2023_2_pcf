package main

import (
	server "AdHub/internal/app"
	log "AdHub/pkg/logger"
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
	l := log.NewLogrusLogger("Error")

	err := server.Parse(configPath)
	if err != nil {
		l.Error("Error: Flag parse " + err.Error())
	}

	s := server.New(server.NewConfig())
	if err := s.Start(); err != nil {
		l.Error("Error: Server start" + err.Error())
	}
}
