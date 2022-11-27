package cache

import (
	"time"

	"github.com/go-redis/redis"
)

var (
	opt = &redis.Options{
		Addr:     "tcache:6379",
		Password: "hellonadya",
		DB:       0,
	}
	tokenTime = 15 * time.Minute
)

type RedisRepo struct {
	client *redis.Client
}

func SetRedis() (*RedisRepo, error) {
	client := redis.NewClient(opt)
	if err := client.Ping().Err(); err != nil {
		return nil, err
	}
	return &RedisRepo{
		client: client,
	}, nil
}

func (r *RedisRepo) SetToken(token string) error {
	if err := r.client.Get(token).Err(); err != nil {
		if err := r.client.Set(token, "ok", tokenTime).Err(); err != nil {
			return err
		}
	}

	return nil
}
func (r *RedisRepo) UnsetToken(token string) error {
	if err := r.client.Del(token).Err(); err != nil {
		return err
	}
	return nil
}
