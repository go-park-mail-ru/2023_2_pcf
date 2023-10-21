package router

import (
	"AdHub/internal/pkg/delivery/middleware"
	"AdHub/internal/pkg/entities"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

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

	httpSwagger.URL("/swagger-docs/swagger.json")
	http.Handle("/swagger/", httpSwagger.WrapHandler)
	http.Handle("/swagger-docs/", http.StripPrefix("/swagger-docs/", http.FileServer(http.Dir("./cmd/apiserver/docs/"))))

	http.Handle("/", ar.router)
}
