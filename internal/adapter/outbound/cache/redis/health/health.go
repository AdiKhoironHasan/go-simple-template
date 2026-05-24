package health

import (
	red "github.com/redis/go-redis/v9"
)

type cacheRepo struct {
	client *red.Client
}

func NewCache(client *red.Client) *cacheRepo {
	return &cacheRepo{
		client: client,
	}
}
