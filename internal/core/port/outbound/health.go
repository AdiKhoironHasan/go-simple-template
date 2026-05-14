package outbound

import (
	"context"
)

//go:generate mockgen -package mocks -source=health.go -destination=mocks/health_mock.go HealthRepository,CacheRepository
type HealthRepository interface {
	CheckHealth(ctx context.Context) error
}

type HealthCacheRepository interface {
	CheckHealth(ctx context.Context) error
}
