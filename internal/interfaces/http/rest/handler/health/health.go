package health

import "go-simple-template/internal/core/ports/service"

type health struct {
	healthService service.HealthService
}

func New(healthService service.HealthService) *health {
	return &health{
		healthService: healthService,
	}
}
