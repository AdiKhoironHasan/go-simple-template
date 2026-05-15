package cache

import (
	"context"
)

//go:generate mockgen -package mocks -source=health.go -destination=mocks/health_mock.go Health
type Health interface {
	CheckHealth(ctx context.Context) error
}
