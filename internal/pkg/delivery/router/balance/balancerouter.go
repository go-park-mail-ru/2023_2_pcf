package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/middleware"

	"AdHub/pkg/logger"

	"github.com/gorilla/mux"
)

type BalanceRouter struct {
	router  *mux.Router
	logger  logger.Logger
	Balance entities.BalanceUseCaseInterface
	Session entities.SessionUseCaseInterface
	User    entities.UserUseCaseInterface
}

func NewBalanceRouter(r *mux.Router, UserUC entities.UserUseCaseInterface, BalanceUC entities.BalanceUseCaseInterface, SessionUC entities.SessionUseCaseInterface, log logger.Logger) *BalanceRouter {
	return &BalanceRouter{
		logger:  log,
		router:  r,
		Balance: BalanceUC,
		Session: SessionUC,
		User:    UserUC,
	}
}

func ConfigureRouter(ur *BalanceRouter) {
	//ur.router.HandleFunc("/ping", PingHandler).Methods("GET", "OPTIONS")
	ur.router.HandleFunc("/balanceadd", ur.BalanceReplenishHandler).Methods("POST", "OPTIONS")
	ur.router.HandleFunc("/balanceget", ur.GetBalanceHandler).Methods("GET", "OPTIONS")

	ur.router.Use(middleware.CORS)
	ur.router.Use(middleware.Auth(ur.Session))
	ur.router.Use(middleware.Logger(ur.logger))
	ur.router.Use(middleware.Recover(ur.logger))
}
