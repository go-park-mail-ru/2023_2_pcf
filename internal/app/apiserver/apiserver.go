package apiserver

import (
	"AdHub/internal/app/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	config *Config
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	s.configureRouter()

	log.Println("INFO: Starting API sever")
	return http.ListenAndServe(s.config.BindAddr, nil)
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/ping", handlers.PingHandler).Methods("GET")

	http.Handle("/", s.router)
}
