package SessionStorage

import "github.com/go-redis/redis/v8"

type SessionStorageInterface interface {
	Store() *redis.Client
	Open() (SessionStorageInterface, error)
	Close()
}
