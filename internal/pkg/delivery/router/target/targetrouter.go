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

func NewTargetRouter(r *mux.Router, TargetUC entities.TargetUseCaseInterface, SessionUC entities.SessionUseCaseInterface, log logger.Logger) *TargetRouter {
	return &TargetRouter{
		logger:  log,
		router:  r,
		Target:  TargetUC,
		Session: SessionUC,
	}
}

func ConfigureRouter(ur *TargetRouter) {
	ur.router.HandleFunc("/targetcreate", ur.CreateTargetHandler).Methods("POST", "OPTIONS")
	ur.router.HandleFunc("/targetedit", ur.UpdateTargetHandler).Methods("POST", "OPTIONS")
	ur.router.HandleFunc("/targetdelete", ur.TargetDeleteHandler).Methods("DELETE", "OPTIONS")
	ur.router.HandleFunc("/targetget", ur.GetTargetHandler).Methods("GET", "OPTIONS")
	ur.router.HandleFunc("/targetlist", ur.TargetListHandler).Methods("GET", "OPTIONS")

	ur.router.Use(middleware.CORS)
	ur.router.Use(middleware.Auth(ur.Session))
	ur.router.Use(middleware.Logger(ur.logger))
	ur.router.Use(middleware.Recover(ur.logger))
}
