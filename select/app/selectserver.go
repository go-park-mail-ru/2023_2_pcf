package server

import (
	"AdHub/pkg/logger"
	api "AdHub/proto/api"
	entities "AdHub/select/pkg/entities"
	"AdHub/select/pkg/usecase/selectuc"
	"context"
	"errors"
	"net"
	"net/http"

	"time"

	"google.golang.org/grpc"

	"google.golang.org/grpc/status"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	grpcRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_requests_total_select",
			Help: "Number of gRPC requests.",
		},
		[]string{"service", "method", "status"},
	)

	grpcRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "grpc_request_duration_seconds_select",
			Help:    "Duration of gRPC requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"service", "method"},
	)
)

func init() {
	prometheus.MustRegister(grpcRequestsTotal)
	prometheus.MustRegister(grpcRequestDuration)
}

func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	startTime := time.Now()
	resp, err := handler(ctx, req)
	duration := time.Since(startTime)

	statusCode := status.Code(err).String()
	service := info.FullMethod // You might want to parse this

	grpcRequestsTotal.WithLabelValues(service, info.FullMethod, statusCode).Inc()
	grpcRequestDuration.WithLabelValues(service, info.FullMethod).Observe(duration.Seconds())

	return resp, err
}

// GRPCServer ...
type GRPCServer struct {
	config   *Config
	SelectUC *selectuc.SelectUseCase
	api.UnimplementedSelectServer
}

func New(config *Config) *GRPCServer {
	return &GRPCServer{
		config: config,
	}
}

func (s *GRPCServer) Start() error {
	log := logger.NewLogrusLogger(s.config.LogLevel)

	serv := grpc.NewServer(grpc.UnaryInterceptor(UnaryServerInterceptor))
	api.RegisterSelectServer(serv, s)
	l, err := net.Listen("tcp", s.config.BindAddr)
	if err != nil {
		log.Error(err.Error())
	}
	s.SelectUC = selectuc.New()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8085", nil)
	}()

	log.Info("Starting Select service on " + s.config.BindAddr)
	return serv.Serve(l)
}

func (s *GRPCServer) Get(ctx context.Context, req *api.SelectRequests) (*api.SelectResponse, error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	targets := make([]*entities.Target, 0, len(req.Requests))

	// Iterate over the requests and parse them
	for _, r := range req.Requests {
		target := &entities.Target{
			Id:        int(r.GetId()),
			Name:      r.GetName(),
			Owner_id:  int(r.GetOwnerId()),
			Gender:    r.GetGender(),
			Min_age:   int(r.GetMinAge()),
			Max_age:   int(r.GetMaxAge()),
			Interests: r.GetInterests(),
			Tags:      r.GetTags(),
			Keys:      r.GetKeys(),
			Regions:   r.GetRegions(),
		}
		targets = append(targets, target)
	}

	pad := &entities.Target{
		Id:        int(req.GetId()),
		Name:      req.GetName(),
		Owner_id:  int(req.GetOwnerId()),
		Gender:    req.GetGender(),
		Min_age:   int(req.GetMinAge()),
		Max_age:   int(req.GetMaxAge()),
		Interests: req.GetInterests(),
		Tags:      req.GetTags(),
		Keys:      req.GetKeys(),
		Regions:   req.GetRegions(),
	}

	ss := s.SelectUC.GetAd(targets, pad)
	return &api.SelectResponse{Id: int64(ss)}, nil
}
