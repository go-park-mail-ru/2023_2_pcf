package router

import (
	"AdHub/internal/interfaces"
	"AdHub/internal/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type MuxRouter struct {
	router *mux.Router
	User   interfaces.UserUseCase
	Ad     interfaces.AdUseCase
}

func NewMuxRouter(UserUC interfaces.UserUseCase, AdUC interfaces.AdUseCase) *MuxRouter {
	return &MuxRouter{
		router: mux.NewRouter(),
		User:   UserUC,
		Ad:     AdUC,
	}
}

func (mr *MuxRouter) ConfigureRouter() {
	mr.router.HandleFunc("/ping", PingHandler).Methods("GET", "OPTIONS")
	mr.router.HandleFunc("/user", mr.UserReadHandler).Methods("GET", "OPTIONS")
	mr.router.HandleFunc("/user", mr.UserCreateHandler).Methods("POST", "OPTIONS")
	mr.router.HandleFunc("/user", mr.UserDeleteHandler).Methods("DELETE", "OPTIONS")
	mr.router.HandleFunc("/ad", mr.AdListHandler).Methods("GET", "OPTIONS")
	mr.router.HandleFunc("/ad", mr.AdCreateHandler).Methods("POST", "OPTIONS")
	mr.router.HandleFunc("/auth", mr.AuthHandler).Methods("POST", "OPTIONS")

	mr.router.Use(middleware.CORSMiddleware)
	mr.router.Use(middleware.RecoverMiddleware)

	http.Handle("/", mr.router)
}
