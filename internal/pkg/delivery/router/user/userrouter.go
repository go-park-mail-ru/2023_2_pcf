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
	Balance entities.BalanceUseCaseInterface
	File    entities.FileUseCaseInterface
}

func NewUserRouter(r *mux.Router, UserUC entities.UserUseCaseInterface, SessionUC entities.SessionUseCaseInterface, FileUC entities.FileUseCaseInterface, BalanceUC entities.BalanceUseCaseInterface, log logger.Logger) *UserRouter {
	return &UserRouter{
		logger:  log,
		router:  r,
		User:    UserUC,
		Session: SessionUC,
		Balance: BalanceUC,
		File:    FileUC,
	}
}

func ConfigureRouter(ur *UserRouter) {
	ur.router.HandleFunc("/ping", PingHandler).Methods("GET", "OPTIONS")
	ur.router.HandleFunc("/user", ur.UserReadHandler).Methods("GET", "OPTIONS")
	ur.router.HandleFunc("/user", ur.UserCreateHandler).Methods("POST", "OPTIONS")
	ur.router.HandleFunc("/userdel", ur.UserDeleteHandler).Methods("POST", "OPTIONS")
	ur.router.HandleFunc("/auth", ur.AuthHandler).Methods("POST", "OPTIONS")
	ur.router.HandleFunc("/useredit", ur.AuthHandler).Methods("POST", "OPTIONS")
	ur.router.HandleFunc("/usergetbytoken", ur.GetUserByTokenHandler).Methods("GET", "OPTIONS")
	ur.router.HandleFunc("/file", ur.FileHandler).Methods("GET", "OPTIONS")

	ur.router.Use(middleware.CORS)
	ur.router.Use(middleware.Auth(ur.Session))
	ur.router.Use(middleware.Logger(ur.logger))
	ur.router.Use(middleware.Recover(ur.logger))
}
