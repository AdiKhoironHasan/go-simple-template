package redis

import (
	"context"
	"fmt"
	"go-simple-template/config"

	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
}

func NewRedis() *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.RedisHost(), config.RedisPort()),
		Password: config.RedisPassword(),
		DB:       config.RedisDB(),
	})
	return &Redis{client: client}
}

type RedisInterface interface {
	Ping() (string, error)
}

func (r *Redis) Ping(ctx context.Context) (string, error) {
	return r.client.WithContext(ctx).Ping().Result()
}
