package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/middleware"

	"AdHub/pkg/logger"

	"github.com/gorilla/mux"
)

type UserRouter struct {
	router  *mux.Router
	logger  logger.Logger
	User    entities.UserUseCaseInterface
	Session entities.SessionUseCaseInterface
	File    entities.FileUseCaseInterface
}

func NewUserRouter(r *mux.Router, UserUC entities.UserUseCaseInterface, SessionUC entities.SessionUseCaseInterface, FileUC entities.FileUseCaseInterface, log logger.Logger) *UserRouter {
	return &UserRouter{
		logger:  log,
		router:  r,
		User:    UserUC,
		Session: SessionUC,
		File:    FileUC,
	}
}

func ConfigureRouter(ur *UserRouter) {
	ur.router.HandleFunc("/ping", PingHandler).Methods("GET", "OPTIONS")
	ur.router.HandleFunc("/user", ur.UserReadHandler).Methods("GET", "OPTIONS")
	ur.router.HandleFunc("/user", ur.UserCreateHandler).Methods("POST", "OPTIONS")
	ur.router.HandleFunc("/user", ur.UserDeleteHandler).Methods("DELETE", "OPTIONS")
	ur.router.HandleFunc("/auth", ur.AuthHandler).Methods("POST", "OPTIONS")
	ur.router.HandleFunc("/useredit", ur.AuthHandler).Methods("POST", "OPTIONS")

	ur.router.Use(middleware.CORS)
	ur.router.Use(middleware.Logger(ur.logger))
	ur.router.Use(middleware.Recover(ur.logger))
}
