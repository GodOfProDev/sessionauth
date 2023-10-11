package storage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage() (*RedisStorage, error) {
	opt, err := redis.ParseURL(os.Getenv("DB_URL"))
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)

	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &RedisStorage{
		client: client,
	}, nil
}

func (s RedisStorage) CreateUser() {

}
