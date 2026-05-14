package health

import (
	"go-simple-template/internal/core/port/inbound"
	"go-simple-template/internal/core/port/outbound"
)

type service struct {
	healthRepo      outbound.HealthRepository
	healthCacheRepo outbound.HealthCacheRepository
}

func New(
	healthRepo outbound.HealthRepository,
	healthCacheRepo outbound.HealthCacheRepository,
) inbound.HealthService {
	return &service{
		healthRepo:      healthRepo,
		healthCacheRepo: healthCacheRepo,
	}
}
