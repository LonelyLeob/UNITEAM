package redisclient

import (
	"context"

	"github.com/go-redis/redis/v9"
)

var (
	opt = &redis.Options{
		Addr:     "tcache:6379",
		Password: "hellonadya",
		DB:       0,
	}
)

type RedisRepo struct {
	client *redis.Client
}

func SetRedis() *RedisRepo {
	return &RedisRepo{
		client: redis.NewClient(opt),
	}
}

func (r *RedisRepo) CompareTokens(token string) error {
	if err := r.client.Get(context.Background(), token).Err(); err != nil {
		return err
	}
	return nil
}
