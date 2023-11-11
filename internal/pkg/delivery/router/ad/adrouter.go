package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/middleware"

	"AdHub/pkg/logger"

	"github.com/gorilla/mux"
)

type AdRouter struct {
	addr    string
	router  *mux.Router
	logger  logger.Logger
	Ad      entities.AdUseCaseInterface
	Session entities.SessionUseCaseInterface
	File    entities.FileUseCaseInterface
	Balance entities.BalanceUseCaseInterface
	User    entities.UserUseCaseInterface
}

func NewAdRouter(addr string, r *mux.Router, AdUC entities.AdUseCaseInterface, UserUC entities.UserUseCaseInterface, SessionUC entities.SessionUseCaseInterface, FileUC entities.FileUseCaseInterface, BalanceUC entities.BalanceUseCaseInterface, log logger.Logger) *AdRouter {
	return &AdRouter{
		addr:    addr,
		router:  r,
		logger:  log,
		Ad:      AdUC,
		Session: SessionUC,
		Balance: BalanceUC,
		File:    FileUC,
		User:    UserUC,
	}
}

func ConfigureRouter(ar *AdRouter) {
	ar.router.HandleFunc("/ad", ar.AdListHandler).Methods("GET", "OPTIONS")
	ar.router.HandleFunc("/ad", ar.AdCreateHandler).Methods("POST", "OPTIONS")
	ar.router.HandleFunc("/adedit", ar.AdUpdateHandler).Methods("POST", "OPTIONS")
	ar.router.HandleFunc("/addelete", ar.AdDeleteHandler).Methods("POST", "OPTIONS")
	ar.router.HandleFunc("/adget", ar.AdGetHandler).Methods("GET", "OPTIONS")
	ar.router.HandleFunc("/addgetamount", ar.AdGetAmountHandler).Methods("GET", "OPTIONS")
	ar.router.HandleFunc("/aduniquelink", ar.AdBannerHandler).Methods("GET", "OPTIONS")

	ar.router.Use(middleware.CORS)
	ar.router.Use(middleware.Auth(ar.Session))
	ar.router.Use(middleware.Logger(ar.logger))
	ar.router.Use(middleware.Recover(ar.logger))
}
