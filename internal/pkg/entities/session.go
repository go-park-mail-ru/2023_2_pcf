package entities

import (
	"AdHub/proto/api"
	"context"

	"google.golang.org/grpc"
)

type Session struct {
	Token  string `json:"token"`
	UserId int    `json:"-"`
}

//go:generate /Users/bincom/go/bin/mockgen -source=session.go -destination=mock_entities/session_mock.go
type SessionUseCaseInterface interface {
	SessionCreate(sr *Session) (*Session, error)
	SessionRead(sr *Session) (*Session, error)
	SessionRemove(sr *Session) error
	SessionContains(sr *Session) (bool, error)
	Auth(ctx context.Context, req *api.AuthRequest, opts ...grpc.CallOption) (*api.AuthResponse, error)
	GetUserId(ctx context.Context, in *api.GetRequest, opts ...grpc.CallOption) (*api.GetResponse, error)
}

type SessionRepoInterface interface {
	Create(s *Session) (*Session, error)
	Remove(s *Session) error
	Read(s *Session) (*Session, error)
	Contains(s *Session) (bool, error)
}
