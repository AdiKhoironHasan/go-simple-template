package inbound

import (
	"context"
	"github.com/adikhoironhasan/go-simple-template/internal/core/domain/entity"
)

//go:generate mockgen -package mocks -source=health.go -destination=mocks/health_mock.go HealthService
type HealthService interface {
	CheckHealth(ctx context.Context, req entity.CheckHealth) (*entity.CheckHealth, error)
}
