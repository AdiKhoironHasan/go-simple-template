package health

import (
	"go-simple-template/internal/core/ports/repository"
	"go-simple-template/internal/core/ports/service"
)

type health struct {
	healthRepo repository.HealthRepository
}

func New(
	healthRepo repository.HealthRepository,
) service.HealthService {
	return &health{
		healthRepo: healthRepo,
	}
}
