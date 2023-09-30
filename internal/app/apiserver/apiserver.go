package apiserver

import (
	"AdHub/internal/app/handlers"
	"AdHub/internal/app/models"
	"AdHub/internal/app/store"
	"log"
	"net/http"

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
	userNew := &models.User{Id: 100, Login: "aSda14", Password: "dasdsa"}
	// попробуй сделтаь new при возврате пользователя
	userNew, err := s.store.User().Create(userNew)
	if err != nil {
		log.Printf("Create")
	}

	// почему-то ругается на строчку Get, возможно create что-то не возращает
	rows, err := s.store.User().Get("aSda14")
	if err != nil {
		log.Printf("sad")
		return err
	}

	defer rows.Close()

	var user struct {
		Id       int
		Login    string
		Password string
	}

	for rows.Next() {
		var id int
		var login, password string

		err := rows.Scan(&user.Id, &user.Login, &user.Password)
		if err != nil {
			log.Printf("sad")
		}
		user.Id = id
		user.Login = login
		user.Password = password
	}

	log.Printf(user.Login)

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
		log.Printf("Database open error")
		return err
	}

	s.store = st
	return nil
}
