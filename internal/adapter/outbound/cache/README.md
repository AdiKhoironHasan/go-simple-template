# Cache Adapters

Outbound adapters for **in-memory caching** and **cache decorators**.

## Structure

```
cache/
├── redis/              # Redis implementations
│   └── <feature>/      # e.g., health/, user/, session/
│       ├── cache.go            # Constructor + struct
│       └── <method>.go         # One file per method
├── memcached/          # Memcached implementations
│   └── <feature>/
│       ├── cache.go
│       └── <method>.go
└── README.md
```

## Implementation Guide

```go
package health

import (
    "go-simple-template/internal/core/port/outbound"
    red "github.com/redis/go-redis/v9"
)

type cacheRepository struct {
    client *red.Client
}

func NewCache(client *red.Client) outbound.HealthCache {
    return &cacheRepository{client: client}
}
```

## Usage Pattern: Decorator

Cache adapters are typically used as **decorators** wrapping a repository:

```go
// In cmd/ composition root:
userRepo := postgresUser.New(pool)                    // primary
userCache := redisUser.NewCache(redisClient, userRepo) // decorator
userSvc := userService.New(userCache)                 // service sees cache
```

## Rules

- Cache adapters implement their own port interface (e.g., `outbound.HealthCache`)
- Never import repository adapters — coupling happens in `cmd/`
- TTL and eviction logic lives here, not in the domain
