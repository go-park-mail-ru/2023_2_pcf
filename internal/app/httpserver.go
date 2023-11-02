package server

import (
	AdRouter "AdHub/internal/pkg/delivery/router/ad"
	BalanceRouter "AdHub/internal/pkg/delivery/router/balance"
	UserRouter "AdHub/internal/pkg/delivery/router/user"
	"AdHub/internal/pkg/repo"
	"AdHub/internal/pkg/usecases/ad"
	"AdHub/internal/pkg/usecases/balance"
	"AdHub/internal/pkg/usecases/file"
	"AdHub/internal/pkg/usecases/session"
	"AdHub/internal/pkg/usecases/user"
	"AdHub/pkg/SessionStorage"
	"AdHub/pkg/db"
	"AdHub/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	config *Config
}

func New(config *Config) *HTTPServer {
	return &HTTPServer{
		config: config,
	}
}

func (s *HTTPServer) Start() error {
	log := logger.NewLogrusLogger(s.config.LogLevel)
	DB := db.New(s.config.DataBase)
	Redis := SessionStorage.New(s.config.Redis_addr, s.config.Redis_password, s.config.Redis_db)

	FileRepo := repo.NewFileRepository(s.config.File_path)
	UserRepo, err := repo.NewUserRepo(DB)
	if err != nil {
		log.Error("User repo error: " + err.Error())
	}
	AdRepo, err := repo.NewAdRepo(DB)
	if err != nil {
		log.Error("Ad repo error: " + err.Error())
	}
	SessionRepo, err := repo.NewSessionRepo(Redis)
	if err != nil {
		log.Error("Ad repo error: " + err.Error())
	}
	BalanceRepo, err := repo.NewBalanceRepo(DB)
	if err != nil {
		log.Error("Balance repo error: " + err.Error())
	}

	FileUC := file.New(FileRepo)
	SessionUC := session.New(SessionRepo)
	AdUC := ad.New(AdRepo)
	UserUC := user.New(UserRepo)
	BalanceUC := balance.New(BalanceRepo)
	rout := mux.NewRouter()

	userrouter := UserRouter.NewUserRouter(rout.PathPrefix("/api/v1").Subrouter(), UserUC, SessionUC, FileUC, log)
	adrouter := AdRouter.NewAdRouter(rout.PathPrefix("/api/v1").Subrouter(), AdUC, SessionUC, FileUC, log)
	balancerouter := BalanceRouter.NewBalanceRouter(rout.PathPrefix("/api/v1").Subrouter(), BalanceUC, SessionUC, log)

	http.Handle("/", rout)

	UserRouter.ConfigureRouter(userrouter)
	AdRouter.ConfigureRouter(adrouter)
	BalanceRouter.ConfigureRouter(balancerouter)

	log.Info("Starting API sever on " + s.config.BindAddr)
	return http.ListenAndServe(s.config.BindAddr, nil)
}
