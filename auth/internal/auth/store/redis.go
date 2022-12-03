package store

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedis(opts ...string) *RedisStore {
	return &RedisStore{
		client: redis.NewClient(&redis.Options{
			Addr:     opts[0],
			Password: opts[1],
		}),
	}
}

func (r *RedisStore) SaveToken(token, id string, exp time.Duration) error {
	if err := r.client.Set(context.Background(), token, id, exp).Err(); err != nil {
		if debug == "True" {
			logrus.Error(err)
		}

		return errUnreachableAction
	}

	return nil
}

func (r *RedisStore) DeleteToken(token string) error {
	if err := r.client.Del(context.Background(), token).Err(); err != nil {
		if debug == "True" {
			logrus.Error(err)
		}

		return errUnreachableAction
	}
	return nil
}
