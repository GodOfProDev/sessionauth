package session

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
	"sessionauth/internal/storage"
)

type RedisSession struct {
	client *redis.Client
	store  storage.Storage
}

func NewRedisSession(store storage.Storage) (*RedisSession, error) {
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
		store:  store,
	}, nil
}

func (s *RedisSession) GenerateSession() {

}

func (s *RedisSession) GetSession() {

}

func (s *RedisSession) LogIn() {

}

func (s *RedisSession) LogOut() {

}
