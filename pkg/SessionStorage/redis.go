package SessionStorage

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

func (r Redis) Open() (SessionStorageInterface, error) {
	r.client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // Адрес Redis-сервера
		Password: "123",            // Пароль, если он установлен
		DB:       0,                // Номер базы данных
	})

	_, err := r.client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r Redis) Close() {
	//todo
}

func (r Redis) Store() *redis.Client {
	return r.client
}
