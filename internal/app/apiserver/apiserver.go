package apiserver

import (
	"AdHub/internal/app/handlers"
	"AdHub/internal/app/store"
	"log"
	"net/http"

	"AdHub/internal/app/models"

	"github.com/gorilla/mux"
)

type APIServer struct {
	config *Config
	router *mux.Router
	store  *store.Store
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
	userNew, err := s.store.User().Add(&models.User{
		Id:       1,
		Login:    "ASD",
		Password: "ASD",
	})

	_, err = s.store.User().Get(userNew.Login)
	if err != nil {
		return err
	}

	//log.Printf(*user.Login, *user.Id)
	log.Printf("INFO: Starting API sever on %s", s.config.BindAddr) // Временный вариант, надо подумать над библиотекой логирования
	return http.ListenAndServe(s.config.BindAddr, nil)
}

// Сюда пишем роуты
func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/ping", handlers.PingHandler).Methods("GET")

	http.Handle("/", s.router)
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}
