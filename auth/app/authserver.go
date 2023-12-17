package server

import (
	entities "AdHub/auth/pkg/entities"
	repo "AdHub/auth/pkg/repo"
	session "AdHub/auth/pkg/usecase/session"
	"AdHub/pkg/SessionStorage"
	"AdHub/pkg/logger"
	api "AdHub/proto/api"
	"context"
	"net"

	"google.golang.org/grpc"
)

// GRPCServer ...
type GRPCServer struct {
	config    *Config
	SessionUC entities.SessionUseCaseInterface
	api.UnimplementedSessionServer
}

func New(config *Config) *GRPCServer {
	return &GRPCServer{
		config: config,
	}
}

func (s *GRPCServer) Start() error {
	log := logger.NewLogrusLogger(s.config.LogLevel)
	Redis := SessionStorage.New(s.config.Redis_addr, s.config.Redis_password, s.config.Redis_db)

	SessionRepo, err := repo.NewSessionRepo(Redis)
	if err != nil {
		log.Error("Ad repo error: " + err.Error())
	}

	s.SessionUC = session.New(SessionRepo)

	serv := grpc.NewServer()
	api.RegisterSessionServer(serv, s)
	l, err := net.Listen("tcp", s.config.BindAddr)
	if err != nil {
		log.Error(err.Error())
	}

	log.Info("Starting Auth sevice on " + s.config.BindAddr)
	return serv.Serve(l)
}

func (s *GRPCServer) Auth(ctx context.Context, req *api.AuthRequest) (*api.AuthResponse, error) {
	user := entities.User{
		Id:          int(req.GetId()),
		Login:       req.GetLogin(),
		Password:    req.GetPassword(),
		FName:       req.GetFName(),
		LName:       req.GetLName(),
		CompanyName: req.GetCompanyName(),
		Avatar:      req.GetAvatar(),
		BalanceId:   int(req.GetBalanceId()),
	}

	ss, err := s.SessionUC.Auth(&user)
	if err != nil {
		return nil, err
	}
	return &api.AuthResponse{Token: ss.Token, UserId: int64(ss.UserId)}, nil
}

func (s *GRPCServer) GetUserId(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	token := req.GetToken()
	id, err := s.SessionUC.GetUserId(token)
	if err != nil {
		return nil, err
	}
	return &api.GetResponse{Id: int64(id)}, nil
}
