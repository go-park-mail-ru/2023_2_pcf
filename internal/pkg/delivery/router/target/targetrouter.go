package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/middleware"

	"AdHub/pkg/logger"

	"github.com/gorilla/mux"
)

type TargetRouter struct {
	router  *mux.Router
	logger  logger.Logger
	Target  entities.TargetUseCaseInterface
	Session entities.SessionUseCaseInterface
}

func NewUserRouter(r *mux.Router, TargetUC entities.TargetUseCaseInterface, SessionUC entities.SessionUseCaseInterface, log logger.Logger) *TargetRouter {
	return &TargetRouter{
		logger:  log,
		router:  r,
		Target:  TargetUC,
		Session: SessionUC,
	}
}

func ConfigureRouter(ur *TargetRouter) {
	ur.router.HandleFunc("/ping", PingHandler).Methods("GET", "OPTIONS")

	ur.router.Use(middleware.CORS)
	ur.router.Use(middleware.Logger(ur.logger))
	ur.router.Use(middleware.Recover(ur.logger))
}
