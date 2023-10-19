package router

import (
	"AdHub/internal/pkg/delivery/middleware"
	"AdHub/internal/pkg/entities"
	"net/http"

	"github.com/gorilla/mux"
)

type UserRouter struct {
	router *mux.Router
	User   entities.UserUseCaseInterface
}

func NewUserRouter(r *mux.Router, UserUC entities.UserUseCaseInterface) *UserRouter {
	return &UserRouter{
		router: r,
		User:   UserUC,
	}
}

func ConfigureRouter(ur *UserRouter) {
	ur.router.HandleFunc("/user", ur.UserReadHandler).Methods("GET", "OPTIONS")
	ur.router.HandleFunc("/user", ur.UserCreateHandler).Methods("POST", "OPTIONS")
	ur.router.HandleFunc("/user", ur.UserDeleteHandler).Methods("DELETE", "OPTIONS")
	ur.router.HandleFunc("/auth", ur.AuthHandler).Methods("POST", "OPTIONS")

	ur.router.Use(middleware.CORS)
	ur.router.Use(middleware.Recover)

	http.Handle("/", ur.router)
}
