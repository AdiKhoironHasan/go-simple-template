package health

import (
	"go-simple-template/internal/core/port/outbound"

	red "github.com/redis/go-redis/v9"
)

type cacheRepository struct {
	client *red.Client
}

func NewCache(client *red.Client) outbound.HealthCacheRepository {
	return &cacheRepository{
		client: client,
	}
}
