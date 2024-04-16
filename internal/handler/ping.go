package handler

import (
	"go-simple-template/internal/dto"
	"go-simple-template/internal/pkg/utils"
	"go-simple-template/pkg/tracer"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *PingHandler) Ping(c echo.Context) error {
	ctx, span := tracer.SpanStart(c.Request().Context(), "Handler.Ping")
	defer span.Finish()

	err := h.service.Ping(ctx)
	if err != nil {
		span.AddError(err)
		response := utils.ApiResponse().
			WithErrors(utils.ErrorResponse("", err.Error())).
			WithCode(http.StatusInternalServerError).
			WithMessage("failed to ping")
		return c.JSON(response.Code, response)
	}

	return c.JSON(http.StatusOK, dto.ApiResponse{
		Code:    http.StatusOK,
		Message: "pong",
	})
}
