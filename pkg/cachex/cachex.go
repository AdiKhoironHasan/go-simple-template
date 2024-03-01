package cachex

import "context"

type Cache struct {
	Client CacheInterface
}

func NewCache(client CacheInterface) *Cache {
	return &Cache{Client: client}
}

type CacheInterface interface {
	Ping(ctx context.Context) (string, error)
}
