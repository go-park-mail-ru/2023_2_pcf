package CsrfStorage

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisCsrf struct {
	client   *redis.Client
	addr     string
	password string
	db       int
}

func NewMock(red *redis.Client) *RedisCsrf {
	return &RedisCsrf{
		client: red,
	}
}

func New(addr string, password string, db int) *RedisCsrf {
	return &RedisCsrf{
		addr:     addr,
		password: password,
		db:       db,
	}
}

func (r *RedisCsrf) Open() (CsrfStorageInterface, error) {
	r.client = redis.NewClient(&redis.Options{
		Addr:     r.addr,
		Password: r.password,
		DB:       r.db,
	})

	_, err := r.client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *RedisCsrf) Close() error {
	if r.client != nil {
		return r.client.Close()
	}
	return nil
}

func (r *RedisCsrf) Store() *redis.Client {
	return r.client
}
