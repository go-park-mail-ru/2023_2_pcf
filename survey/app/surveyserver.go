package server

import (
	"AdHub/pkg/db"
	"AdHub/pkg/logger"
	SurveyRouter "AdHub/survey/pkg/delivery/router/survey"
	"AdHub/survey/pkg/repo"
	"AdHub/survey/pkg/usecase/rate"
	"AdHub/survey/pkg/usecase/survey"
	"net/http"

	"google.golang.org/grpc"

	"AdHub/proto/api"

	"github.com/gorilla/mux"
)

// HTTPServer ...
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
	DB := db.New(s.config.Db)

	authconn, err := grpc.Dial(s.config.AuthBindAddr, grpc.WithInsecure())
	if err != nil {
		log.Error("Auth Micro Service: " + err.Error())
	}

	Survey, err := repo.NewSurveyRepo(DB)
	if err != nil {
		log.Error("Survey repo error: " + err.Error())
	}

	Rate, err := repo.NewRateRepo(DB)
	if err != nil {
		log.Error("Rate repo error: " + err.Error())
	}

	SurveyUC := survey.New(Survey)
	RateUC := rate.New(Rate)
	rout := mux.NewRouter()

	SessionMS := api.NewSessionClient(authconn)

	surveyrouter := SurveyRouter.NewSurveyRouter(rout.PathPrefix("/api/v1/survey").Subrouter(), SurveyUC, SessionMS, RateUC, log)

	http.Handle("/", rout)

	SurveyRouter.ConfigureRouter(surveyrouter)

	log.Info("Starting API sever on " + s.config.BindAddr)
	return http.ListenAndServe(s.config.BindAddr, nil)
}
