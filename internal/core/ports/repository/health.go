package repository

import (
	"context"
)

//go:generate mockgen -package mocks -source=health.go -destination=mocks/health_mock.go HealthRepository
type HealthRepository interface {
	CheckHealth(ctx context.Context) error
}
