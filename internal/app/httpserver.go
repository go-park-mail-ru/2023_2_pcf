package server

import (
	AdRouter "AdHub/internal/pkg/delivery/router/ad"
	UserRouter "AdHub/internal/pkg/delivery/router/user"
	"AdHub/internal/pkg/repo"
	"AdHub/internal/pkg/usecases/ad"
	"AdHub/internal/pkg/usecases/user"
	"AdHub/pkg/db"
	"AdHub/pkg/logger"
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
	log := logger.NewLogrusLogger()
	DB := db.New()

	UserRepo, err := repo.NewUserRepo(DB)
	if err != nil {
		log.Error("User repo error: " + err.Error())
	}
	AdRepo, err := repo.NewAdRepo(DB)
	if err != nil {
		log.Error("Ad repo error: " + err.Error())
	}

	AdUC := ad.New(AdRepo)
	UserUC := user.New(UserRepo)
	rout := mux.NewRouter()

	userrouter := UserRouter.NewUserRouter(rout.PathPrefix("/api/v1").Subrouter(), UserUC, log)
	adrouter := AdRouter.NewAdRouter(rout.PathPrefix("/api/v1").Subrouter(), AdUC, log)

	http.Handle("/", rout)

	UserRouter.ConfigureRouter(userrouter)
	AdRouter.ConfigureRouter(adrouter)

	log.Info("Starting API sever on " + s.config.BindAddr)
	return http.ListenAndServe(s.config.BindAddr, nil)
}
