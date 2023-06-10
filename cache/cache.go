package cache

type Cache struct {
	Client CacheInterface
}

func NewCache(client CacheInterface) *Cache {
	return &Cache{Client: client}
}

type CacheInterface interface {
	Ping() (string, error)
}
