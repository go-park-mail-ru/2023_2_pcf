package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/middleware"

	"AdHub/pkg/logger"

	"github.com/gorilla/mux"
)

type AdRouter struct {
	router  *mux.Router
	logger  logger.Logger
	Ad      entities.AdUseCaseInterface
	Session entities.SessionUseCaseInterface
}

func NewAdRouter(r *mux.Router, AdUC entities.AdUseCaseInterface, SessionUC entities.SessionUseCaseInterface, log logger.Logger) *AdRouter {
	return &AdRouter{
		router:  r,
		logger:  log,
		Ad:      AdUC,
		Session: SessionUC,
	}
}

func ConfigureRouter(ar *AdRouter) {
	ar.router.HandleFunc("/ad", ar.AdListHandler).Methods("GET", "OPTIONS")
	ar.router.HandleFunc("/ad", ar.AdCreateHandler).Methods("POST", "OPTIONS")
	ar.router.HandleFunc("/adedit", ar.AdUpdateHandler).Methods("POST", "OPTIONS")
	ar.router.HandleFunc("/addelete", ar.AdDeleteHandler).Methods("DELETE", "OPTIONS")
	ar.router.HandleFunc("/addget/{adID}", ar.AdDeleteHandler).Methods("GET", "OPTIONS")

	ar.router.Use(middleware.CORS)
	ar.router.Use(middleware.Logger(ar.logger))
	ar.router.Use(middleware.Recover(ar.logger))
}
