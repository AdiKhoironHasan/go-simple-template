package health

import (
	"go-simple-template/internal/core/port/inbound"
	"go-simple-template/internal/core/port/outbound/cache"
	"go-simple-template/internal/core/port/outbound/repository"
)

type service struct {
	healthRepo  repository.Health
	healthCache cache.Health
}

func New(
	healthRepo repository.Health,
	healthCache cache.Health,
) inbound.HealthService {
	return &service{
		healthRepo:  healthRepo,
		healthCache: healthCache,
	}
}
