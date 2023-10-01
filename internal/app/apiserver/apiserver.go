package apiserver

import (
	"AdHub/internal/app/store"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	config *Config
	router *mux.Router
	Store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureStore(); err != nil {
		return err
	}

	s.configureRouter()

	log.Printf("INFO: Starting API sever on %s", s.config.BindAddr) // Временный вариант, надо подумать над библиотекой логирования
	return http.ListenAndServe(s.config.BindAddr, nil)
}

// Сюда пишем роуты
func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/ping", PingHandler).Methods("GET")
	s.router.HandleFunc("/user", s.UserReadHandler).Methods("GET")
	s.router.HandleFunc("/user", s.UserCreateHandler).Methods("POST")
	//s.router.HandleFunc("/user/{user_id:[0-9]+}", handlers.UserUpdateHandler).Methods("POST")
	s.router.HandleFunc("/user", s.UserDeleteHandler).Methods("DELETE")
	//s.router.HandleFunc("/ad/", handlers.AdHandler).Methods("GET")
	//s.router.HandleFunc("/ad/{ad_id:[0-9]+}", handlers.AdReadHandler).Methods("GET")
	//s.router.HandleFunc("/ad", handlers.AdCreateHandler).Methods("POST")
	//s.router.HandleFunc("/ad/{ad_id:[0-9]+}", handlers.AdUpdateHandler).Methods("POST")
	//s.router.HandleFunc("/ad/{ad_id:[0-9]+}", handlers.AdDeleteHandler).Methods("DELETE")

	http.Handle("/", s.router)
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		log.Printf("Database open error")
		return err
	}

	s.Store = st
	return nil
}
