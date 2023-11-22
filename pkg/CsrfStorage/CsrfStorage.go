package CsrfStorage

import "github.com/go-redis/redis/v8"

type CsrfStorageInterface interface {
	Store() *redis.Client
	Open() (CsrfStorageInterface, error)
	Close() error
}
