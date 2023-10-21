package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/middleware"

	"AdHub/pkg/logger"

	"github.com/gorilla/mux"
)

type AdRouter struct {
	router *mux.Router
	logger logger.Logger
	Ad     entities.AdUseCaseInterface
}

func NewAdRouter(r *mux.Router, AdUC entities.AdUseCaseInterface, log logger.Logger) *AdRouter {
	return &AdRouter{
		router: r,
		logger: log,
		Ad:     AdUC,
	}
}

func ConfigureRouter(ar *AdRouter) {
	ar.router.HandleFunc("/ad", ar.AdListHandler).Methods("GET", "OPTIONS")
	ar.router.HandleFunc("/ad", ar.AdCreateHandler).Methods("POST", "OPTIONS")

	ar.router.Use(middleware.CORS)
	ar.router.Use(middleware.Logger(ar.logger))
	ar.router.Use(middleware.Recover(ar.logger))
}
