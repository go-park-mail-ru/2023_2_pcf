package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/middleware"

	"AdHub/pkg/logger"

	"github.com/gorilla/mux"
)

type PublicRouter struct {
	router *mux.Router
	logger logger.Logger
	Ad     entities.AdUseCaseInterface
}

func NewPublicRouter(r *mux.Router, AdUC entities.AdUseCaseInterface, log logger.Logger) *PublicRouter {
	return &PublicRouter{
		router: r,
		logger: log,
		Ad:     AdUC,
	}
}

func ConfigureRouter(ar *PublicRouter) {
	ar.router.HandleFunc("/redirect", ar.RedirectHandler).Methods("GET", "OPTIONS")

	ar.router.Use(middleware.Pub_CORS)
	ar.router.Use(middleware.Logger(ar.logger))
	ar.router.Use(middleware.Recover(ar.logger))
}
