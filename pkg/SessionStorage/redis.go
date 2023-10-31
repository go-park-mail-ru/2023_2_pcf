package SessionStorage

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client   *redis.Client
	addr     string
	password string
	db       int
}

func NewMock(red *redis.Client) *Redis {
	return &Redis{
		client: red,
	}
}

func New(addr string, password string, db int) *Redis {
	return &Redis{
		addr:     addr,
		password: password,
		db:       db,
	}
}

func (r Redis) Open() (SessionStorageInterface, error) {
	r.client = redis.NewClient(&redis.Options{
		Addr: r.addr, // Адрес Redis-сервера
		DB:   r.db,   // Номер базы данных
	})

	_, err := r.client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r Redis) Close() error {
	return nil
}

func (r Redis) Store() *redis.Client {
	return r.client
}
