package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/logger"
	"AdHub/pkg/middleware"
	"AdHub/proto/api"

	"github.com/gorilla/mux"
)

type PublicRouter struct {
	router   *mux.Router
	logger   logger.Logger
	Ad       entities.AdUseCaseInterface
	SelectUC api.SelectClient
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
	ar.router.HandleFunc("/aduniquelink", ar.AdBannerHandler).Methods("GET", "OPTIONS")

	ar.router.Use(middleware.Pub_CORS)
	ar.router.Use(middleware.Logger(ar.logger))
	ar.router.Use(middleware.Recover(ar.logger))
}
