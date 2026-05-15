package health

import (
	"go-simple-template/internal/core/port/outbound/cache"

	red "github.com/redis/go-redis/v9"
)

type cacheRepo struct {
	client *red.Client
}

func NewCache(client *red.Client) cache.Health {
	return &cacheRepo{
		client: client,
	}
}
