package redis

import (
	"fmt"
	"go-simple-template/config"

	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(cfg *config.Config) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.RedisHost, cfg.Redis.RedisPort),
		Password: cfg.Redis.RedisPassword,
		DB:       cfg.Redis.RedisDB,
	})
	return &Redis{client: client}
}

type RedisInterface interface {
	Ping() (string, error)
}

func (r *Redis) Ping() (string, error) {
	return r.client.Ping().Result()
}
