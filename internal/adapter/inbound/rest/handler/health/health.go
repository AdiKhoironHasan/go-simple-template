package health

import "github.com/adikhoironhasan/go-simple-template/internal/core/port/inbound"

type handler struct {
	healthService inbound.HealthService
}

func New(healthService inbound.HealthService) *handler {
	return &handler{
		healthService: healthService,
	}
}
