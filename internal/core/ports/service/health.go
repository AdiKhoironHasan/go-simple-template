package service

import "context"

//go:generate mockgen -package mocks -source=health.go -destination=mocks/health_mock.go HealthService
type HealthService interface {
	CheckHealth(ctx context.Context) error
}
