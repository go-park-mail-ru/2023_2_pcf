package server

import (
	AdRouter "AdHub/internal/pkg/delivery/router/ad"
	UserRouter "AdHub/internal/pkg/delivery/router/user"
	"AdHub/internal/pkg/repo"
	"AdHub/internal/pkg/usecases/ad"
	"AdHub/internal/pkg/usecases/user"
	"AdHub/pkg/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	DB := db.New()
	UserRepo, err := repo.NewUserRepo(DB)
	if err != nil {

	}
	AdRepo, err := repo.NewAdRepo(DB)
	if err != nil {

	}
	AdUC := ad.New(AdRepo)
	UserUC := user.New(UserRepo)
	rout := mux.NewRouter()

	userrouter := UserRouter.NewUserRouter(rout, UserUC)
	adrouter := AdRouter.NewAdRouter(rout, AdUC)

	UserRouter.ConfigureRouter(userrouter)
	AdRouter.ConfigureRouter(adrouter)

	log.Printf("INFO: Starting API sever on %s", s.config.BindAddr) // Временный вариант, надо подумать над библиотекой логирования
	return http.ListenAndServe(s.config.BindAddr, nil)
}
