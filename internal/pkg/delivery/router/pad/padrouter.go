package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/middleware"

	"AdHub/pkg/logger"

	"AdHub/proto/api"

	"github.com/gorilla/mux"
)

type PadRouter struct {
	addr    string
	router  *mux.Router
	logger  logger.Logger
	Ad      entities.AdUseCaseInterface
	Csrf    entities.CsrfUseCaseInterface
	Session api.SessionClient
	File    entities.FileUseCaseInterface
	Balance entities.BalanceUseCaseInterface
	User    entities.UserUseCaseInterface
	Pad     entities.PadUseCaseInterface
}

func NewPadRouter(addr string, r *mux.Router, AdUC entities.AdUseCaseInterface, UserUC entities.UserUseCaseInterface, CsrfUC entities.CsrfUseCaseInterface, SessionUC api.SessionClient, FileUC entities.FileUseCaseInterface, BalanceUC entities.BalanceUseCaseInterface, PadUC entities.PadUseCaseInterface, log logger.Logger) *PadRouter {
	return &PadRouter{
		addr:    addr,
		router:  r,
		logger:  log,
		Ad:      AdUC,
		Csrf:    CsrfUC,
		Session: SessionUC,
		Balance: BalanceUC,
		File:    FileUC,
		User:    UserUC,
		Pad:     PadUC,
	}
}

func ConfigureRouter(ar *PadRouter) {
	ar.router.HandleFunc("/pad", ar.PadListHandler).Methods("GET", "OPTIONS")
	ar.router.HandleFunc("/pad", ar.PadCreateHandler).Methods("POST", "OPTIONS")
	ar.router.HandleFunc("/padedit", ar.PadUpdateHandler).Methods("POST", "OPTIONS")
	ar.router.HandleFunc("/paddelete", ar.PadDeleteHandler).Methods("POST", "OPTIONS")

	ar.router.Use(middleware.CORS)
	ar.router.Use(middleware.Auth(ar.Session))
	ar.router.Use(middleware.Logger(ar.logger))
	ar.router.Use(middleware.Recover(ar.logger))
}
