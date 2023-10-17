package server

import (
	"AdHub/internal/app/frameworks/store"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	config *Config
	router interfaces.router
	Store  interfaces.db
}

func New(config *Config) *HTTPServer {
	return &HTTPServer{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *HTTPServer) Start() error {
	if err := s.configureStore(); err != nil {
		return err
	}

	s.configureRouter()

	log.Printf("INFO: Starting API sever on %s", s.config.BindAddr) // Временный вариант, надо подумать над библиотекой логирования
	return http.ListenAndServe(s.config.BindAddr, nil)
}

// Сюда пишем роуты
func (s *HTTPServer) configureRouter() {
	s.router.HandleFunc("/ping", PingHandler).Methods("GET")
	s.router.HandleFunc("/user", s.UserReadHandler).Methods("GET")
	s.router.HandleFunc("/user", s.UserCreateHandler).Methods("POST", "OPTIONS")
	//s.router.HandleFunc("/user/{user_id:[0-9]+}", handlers.UserUpdateHandler).Methods("POST")
	s.router.HandleFunc("/user", s.UserDeleteHandler).Methods("DELETE")
	s.router.HandleFunc("/ad", s.AdListHandler).Methods("GET")
	//s.router.HandleFunc("/ad/{ad_id:[0-9]+}", handlers.AdReadHandler).Methods("GET")
	s.router.HandleFunc("/ad", s.AdCreateHandler).Methods("POST")
	//s.router.HandleFunc("/ad/{ad_id:[0-9]+}", handlers.AdUpdateHandler).Methods("POST")
	//s.router.HandleFunc("/ad/{ad_id:[0-9]+}", handlers.AdDeleteHandler).Methods("DELETE")
	s.router.HandleFunc("/auth", s.AuthHandler).Methods("POST", "OPTIONS")

	http.Handle("/", s.router)
}

func (s *HTTPServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		log.Printf("Database open error")
		return err
	}

	s.Store = st
	return nil
}
