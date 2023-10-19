package router

import (
	"AdHub/internal/pkg/delivery/middleware"
	"AdHub/internal/pkg/entities"
	"net/http"

	"github.com/gorilla/mux"
)

type AdRouter struct {
	router *mux.Router
	Ad     entities.AdUseCaseInterface
}

func NewAdRouter(r *mux.Router, AdUC entities.AdUseCaseInterface) *AdRouter {
	return &AdRouter{
		router: r,
		Ad:     AdUC,
	}
}

func ConfigureRouter(ar *AdRouter) {
	ar.router.HandleFunc("/ad", ar.AdListHandler).Methods("GET", "OPTIONS")
	ar.router.HandleFunc("/ad", ar.AdCreateHandler).Methods("POST", "OPTIONS")

	ar.router.Use(middleware.CORS)
	ar.router.Use(middleware.Recover)

	http.Handle("/", ar.router)
}
