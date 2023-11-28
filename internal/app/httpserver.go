package server

import (
	AdRouter "AdHub/internal/pkg/delivery/router/ad"
	BalanceRouter "AdHub/internal/pkg/delivery/router/balance"
	PadRouter "AdHub/internal/pkg/delivery/router/pad"
	PublicRouter "AdHub/internal/pkg/delivery/router/public"
	TargetRouter "AdHub/internal/pkg/delivery/router/target"
	UserRouter "AdHub/internal/pkg/delivery/router/user"
	"AdHub/internal/pkg/repo"
	"AdHub/internal/pkg/usecases/ad"
	"AdHub/internal/pkg/usecases/balance"
	"AdHub/internal/pkg/usecases/csrf"
	"AdHub/internal/pkg/usecases/file"
	"AdHub/internal/pkg/usecases/pad"
	"AdHub/internal/pkg/usecases/target"
	"AdHub/internal/pkg/usecases/ulink"
	"AdHub/internal/pkg/usecases/user"
	"AdHub/pkg/CsrfStorage"
	"AdHub/pkg/SessionStorage"
	"AdHub/pkg/db"
	"AdHub/pkg/logger"
	"AdHub/pkg/middleware"
	"AdHub/proto/api"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
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
	Redis := SessionStorage.New(s.config.Redis_addr, s.config.Redis_password, s.config.Redis_db_ul)
	Redis_CSRF := CsrfStorage.New(s.config.Redis_addr, s.config.Redis_password, s.config.Redis_db_csrf)

	ULinkRepo, err := repo.NewULinkRepo(Redis)
	if err != nil {
		log.Error("Ad repo error: " + err.Error())
	}
	ULinkUC := ulink.New(ULinkRepo)

	CSRFRepo, err := repo.NewCsrfRepo(Redis_CSRF)
	if err != nil {
		log.Error("CSRF repo error: " + err.Error())
	}
	CsrfUC := csrf.New(CSRFRepo)

	FileRepo := repo.NewFileRepository(s.config.File_path)
	UserRepo, err := repo.NewUserRepo(DB)
	if err != nil {
		log.Error("User repo error: " + err.Error())
	}
	AdRepo, err := repo.NewAdRepo(DB)
	if err != nil {
		log.Error("Ad repo error: " + err.Error())
	}
	BalanceRepo, err := repo.NewBalanceRepo(DB)
	if err != nil {
		log.Error("Balance repo error: " + err.Error())
	}
	TargetRepo, err := repo.NewTargetRepo(DB)
	if err != nil {
		log.Error("Target repo error: " + err.Error())
	}
	PadRepo, err := repo.NewPadRepo(DB)
	if err != nil {
		log.Error("Pad repo error: " + err.Error())
	}

	authconn, err := grpc.Dial(s.config.AuthBindAddr, grpc.WithInsecure())
	if err != nil {
		log.Error("Auth Micro Service: " + err.Error())
	}

	selectconn, err := grpc.Dial(s.config.SelectBindAddr, grpc.WithInsecure())

	FileUC := file.New(FileRepo)
	SessionMS := api.NewSessionClient(authconn)
	SelectMS := api.NewSelectClient(selectconn)
	AdUC := ad.New(AdRepo)
	UserUC := user.New(UserRepo)
	BalanceUC := balance.New(BalanceRepo)
	TargetUC := target.New(TargetRepo)
	PadUC := pad.New(PadRepo)
	rout := mux.NewRouter()
	rout.Use(middleware.MetricsMiddleware)

	userrouter := UserRouter.NewUserRouter(rout.PathPrefix("/api/v1").Subrouter(), UserUC, SessionMS, CsrfUC, FileUC, BalanceUC, log)
	adrouter := AdRouter.NewAdRouter(s.config.BindAddr, rout.PathPrefix("/api/v1").Subrouter(), AdUC, UserUC, CsrfUC, SessionMS, FileUC, BalanceUC, log)
	padrouter := PadRouter.NewPadRouter(s.config.BindAddr, rout.PathPrefix("/api/v1").Subrouter(), AdUC, UserUC, CsrfUC, SessionMS, FileUC, BalanceUC, PadUC, log)
	balancerouter := BalanceRouter.NewBalanceRouter(rout.PathPrefix("/api/v1").Subrouter(), UserUC, BalanceUC, CsrfUC, SessionMS, log)
	targetrouter := TargetRouter.NewTargetRouter(rout.PathPrefix("/api/v1").Subrouter(), TargetUC, CsrfUC, SessionMS, log)
	publicRouter := PublicRouter.NewPublicRouter(rout.PathPrefix("/api/v1").Subrouter(), s.config.BindAddr, UserUC, ULinkUC, AdUC, TargetUC, PadUC, SelectMS, log)
	http.Handle("/", rout)
	rout.Handle("/metrics", promhttp.Handler())

	UserRouter.ConfigureRouter(userrouter)
	AdRouter.ConfigureRouter(adrouter)
	BalanceRouter.ConfigureRouter(balancerouter)
	TargetRouter.ConfigureRouter(targetrouter)
	PadRouter.ConfigureRouter(padrouter)
	PublicRouter.ConfigureRouter(publicRouter)

	log.Info("Starting API sever on " + s.config.BindAddr)
	return http.ListenAndServe(s.config.BindAddr, nil)
}
