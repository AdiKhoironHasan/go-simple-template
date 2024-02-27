package handler

import (
	"go-simple-template/internal/dto"
	"go-simple-template/pkg/tracer"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Ping(c echo.Context) error {
	ctx, span := tracer.SpanStart(c.Request().Context(), "Handler.Ping")
	defer span.Finish()

	err := h.service.Ping(ctx)
	if err != nil {
		span.AddError(err)
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
