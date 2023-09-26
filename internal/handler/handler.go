package handler

import (
	"go-simple-template/internal/service"
)

type Handler struct {
	service service.ServiceInterface
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) WithService(service service.ServiceInterface) *Handler {
	h.service = service
	return h
}
