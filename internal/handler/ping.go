package handler

import (
	"go-simple-template/pkg/api/rest"
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
		response := rest.ApiResponse().
			WithErrors(rest.ErrorResponse("", err.Error())).
			WithTraceId(span.TraceId()).
			WithCode(http.StatusInternalServerError)
		return c.JSON(response.Meta.Code, response)
	}

	response := rest.ApiResponse()
	return c.JSON(response.Meta.Code, response)
}
