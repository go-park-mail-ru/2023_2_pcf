package SessionStorage

import "github.com/go-redis/redis/v8"

//go:generate /Users/bincom/go/bin/mockgen -source=SessionStorage.go -destination=mock_SessionStorage/mock.go
type SessionStorageInterface interface {
	Store() *redis.Client
	Open() (SessionStorageInterface, error)
	Close() error
}
