package session

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

type RedisSession struct {
	client *redis.Client
}

func NewRedisSession() (*RedisSession, error) {
	opt, err := redis.ParseURL(os.Getenv("REDIS_DB_URL"))
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)

	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &RedisSession{
		client: client,
	}, nil
}

func (s RedisSession) CreateSession() {

}
