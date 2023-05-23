package handler

import (
	"go-simple-template/service"
)

type Handler struct {
	service service.ServiceInterface
}

func NewHandler(service service.ServiceInterface) *Handler {
	return &Handler{
		service: service,
	}
}
