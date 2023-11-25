package router

import (
	"AdHub/pkg/middleware"
	"AdHub/survey/pkg/entities"

	"AdHub/pkg/logger"

	"AdHub/proto/api"

	"github.com/gorilla/mux"
)

type SurveyRouter struct {
	router  *mux.Router
	logger  logger.Logger
	Survey  entities.SurveyUseCaseInterface
	Session api.SessionClient
	Rate    entities.RateUseCaseInterface
}

func NewSurveyRouter(r *mux.Router, SurveyUC entities.SurveyUseCaseInterface, SessionUC api.SessionClient, RateUC entities.RateUseCaseInterface, log logger.Logger) *SurveyRouter {
	return &SurveyRouter{
		logger:  log,
		router:  r,
		Survey:  SurveyUC,
		Session: SessionUC,
		Rate:    RateUC,
	}
}

func ConfigureRouter(ur *SurveyRouter) {
	ur.router.HandleFunc("/get", ur.GetSurvey).Methods("GET", "OPTIONS")
	ur.router.HandleFunc("/stat", ur.GetStat).Methods("GET", "OPTIONS")
	ur.router.HandleFunc("/rate", ur.RateCreateHandler).Methods("POST", "OPTIONS")
	ur.router.HandleFunc("/survey", ur.SurveyCreateHandler).Methods("POST", "OPTIONS")

	ur.router.Use(middleware.CORS)
	ur.router.Use(middleware.Auth(ur.Session))
	ur.router.Use(middleware.Logger(ur.logger))
	ur.router.Use(middleware.Recover(ur.logger))
}
