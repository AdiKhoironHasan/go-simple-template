package handler

import (
	"go-simple-template/service"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service service.ServiceInterface
}

func NewHandler(service service.ServiceInterface) *Handler {
	return &Handler{
		service: service,
	}
}

// handlers
func (h *Handler) Ping(c echo.Context) error {
	err := h.service.Ping()
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, "pong")
}
