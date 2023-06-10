package handler

import (
	"go-simple-template/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Ping(c echo.Context) error {
	err := h.service.Ping()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.ApiResponse{
		Code:    http.StatusOK,
		Message: "pong",
	})
}
