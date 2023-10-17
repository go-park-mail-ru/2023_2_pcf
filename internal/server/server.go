package server

import (
	"log"
	"net/http"

	"AdHub/internal/router"
	"AdHub/pkg/postgres"

	"AdHub/internal/usecases/ad"
	"AdHub/internal/usecases/repo"
	"AdHub/internal/usecases/user"
)

type HTTPServer struct {
	config *Config
}

func New(config *Config) *HTTPServer {
	return &HTTPServer{
		config: config,
	}
}

func (s *HTTPServer) Start() error {
	DB := postgres.New()
	UserRepo, err := repo.NewUserRepo(DB)
	if err != nil {

	}
	AdRepo, err := repo.NewAdRepo(DB)
	if err != nil {

	}
	AdUC := ad.New(AdRepo)
	UserUc := user.New(UserRepo)
	rout := router.NewMuxRouter(UserUc, AdUC)

	rout.ConfigureRouter()

	log.Printf("INFO: Starting API sever on %s", s.config.BindAddr) // Временный вариант, надо подумать над библиотекой логирования
	return http.ListenAndServe(s.config.BindAddr, nil)
}
