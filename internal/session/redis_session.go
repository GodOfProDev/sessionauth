package session

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"os"
	"time"
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

func (s *RedisSession) GenerateSession(userId string) (string, error) {
	sessionId := uuid.NewString()
	if err := s.client.SetEX(context.Background(), sessionId, userId, 24*time.Hour); err != nil {
		return "", err.Err()
	}

	return sessionId, nil
}

func (s *RedisSession) GetUserBySession(session string) (string, error) {
	userId, err := s.client.Get(context.Background(), session).Result()
	if err != nil {
		return "", err
	}

	return userId, nil
}

func (s *RedisSession) DeleteSession(session string) error {
	return s.client.Del(context.Background(), session).Err()
}
