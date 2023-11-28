package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/middleware"

	"AdHub/pkg/logger"

	"AdHub/proto/api"

	"github.com/gorilla/mux"
)

type TargetRouter struct {
	router  *mux.Router
	logger  logger.Logger
	Target  entities.TargetUseCaseInterface
	Csrf    entities.CsrfUseCaseInterface
	Session api.SessionClient
}

func NewTargetRouter(r *mux.Router, TargetUC entities.TargetUseCaseInterface, CsrfUC entities.CsrfUseCaseInterface, SessionUC api.SessionClient, log logger.Logger) *TargetRouter {
	return &TargetRouter{
		logger:  log,
		router:  r,
		Target:  TargetUC,
		Csrf:    CsrfUC,
		Session: SessionUC,
	}
}

func ConfigureRouter(ur *TargetRouter) {
	ur.router.HandleFunc("/targetcreate", ur.CreateTargetHandler).Methods("POST", "OPTIONS")
	ur.router.HandleFunc("/targetedit", ur.UpdateTargetHandler).Methods("POST", "OPTIONS")
	ur.router.HandleFunc("/targetdelete", ur.TargetDeleteHandler).Methods("POST", "OPTIONS")
	ur.router.HandleFunc("/targetget", ur.GetTargetHandler).Methods("GET", "OPTIONS")
	ur.router.HandleFunc("/targetlist", ur.TargetListHandler).Methods("GET", "OPTIONS")

	ur.router.Use(middleware.CORS)
	ur.router.Use(middleware.Auth(ur.Session, ur.Csrf))
	ur.router.Use(middleware.Logger(ur.logger))
	ur.router.Use(middleware.Recover(ur.logger))
}
