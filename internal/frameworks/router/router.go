package router

import (
	"AdHub/internal/app/interfaces"
	"net/http"

	"github.com/gorilla/mux"
)

type MuxRouter struct {
	router *mux.Router
}

func NewMuxRouter() interfaces.Router {
	return &MuxRouter{
		router: mux.NewRouter(),
	}
}

func (mr *MuxRouter) HandleFunc(path string, handler func(http.ResponseWriter, *http.Request)) {
	mr.router.HandleFunc(path, handler)
}

func (mr *MuxRouter) ConfigureRouter() {
	mr.router.HandleFunc("/ping", PingHandler).Methods("GET")
	mr.router.HandleFunc("/user", s.UserReadHandler).Methods("GET")
	mr.router.HandleFunc("/user", s.UserCreateHandler).Methods("POST", "OPTIONS")
	mr.router.HandleFunc("/user", s.UserDeleteHandler).Methods("DELETE")
	mr.router.HandleFunc("/ad", s.AdListHandler).Methods("GET")
	mr.router.HandleFunc("/ad", s.AdCreateHandler).Methods("POST")
	mr.router.HandleFunc("/auth", s.AuthHandler).Methods("POST", "OPTIONS")

	http.Handle("/", mr.router)
}
