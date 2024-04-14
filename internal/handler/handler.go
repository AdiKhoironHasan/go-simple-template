package handler

import (
	"go-simple-template/internal/service"
)

type PingHandler struct {
	service service.PingService
}

func NewPing(service service.PingService) *PingHandler {
	return &PingHandler{
		service: service,
	}
}
