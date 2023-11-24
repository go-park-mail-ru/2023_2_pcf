package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/pkg/logger"
	"AdHub/pkg/middleware"
	"AdHub/proto/api"

	"github.com/gorilla/mux"
)

type PublicRouter struct {
	addr     string
	router   *mux.Router
	logger   logger.Logger
	ULink    entities.ULinkUseCaseInterface
	Ad       entities.AdUseCaseInterface
	Target   entities.TargetUseCaseInterface
	Pad      entities.PadUseCaseInterface
	SelectUC api.SelectClient
}

func NewPublicRouter(r *mux.Router, addr string, ULinkUC entities.ULinkUseCaseInterface, AdUC entities.AdUseCaseInterface, TargetUC entities.TargetUseCaseInterface, PadUC entities.PadUseCaseInterface, Select api.SelectClient, log logger.Logger) *PublicRouter {
	return &PublicRouter{
		addr:     addr,
		router:   r,
		logger:   log,
		ULink:    ULinkUC,
		Ad:       AdUC,
		Target:   TargetUC,
		Pad:      PadUC,
		SelectUC: Select,
	}
}

func ConfigureRouter(ar *PublicRouter) {
	ar.router.HandleFunc("/redirect", ar.RedirectHandler).Methods("GET", "OPTIONS")
	ar.router.HandleFunc("/getad", ar.GetBanner).Methods("GET", "OPTIONS")

	ar.router.Use(middleware.Pub_CORS)
	ar.router.Use(middleware.Logger(ar.logger))
	ar.router.Use(middleware.Recover(ar.logger))
}
