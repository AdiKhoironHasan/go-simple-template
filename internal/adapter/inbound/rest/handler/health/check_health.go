package health

import (
	"log/slog"
	"net/http"

	"go-simple-template/internal/adapter/inbound/rest/dto"
	"go-simple-template/internal/adapter/inbound/rest/utils"
	"go-simple-template/internal/pkg/consts"

	"github.com/labstack/echo/v4"
)

func (h *handler) CheckHealth(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		request dto.CheckHealthRequest
	)

	response, err := dto.ValidateRequest(ctx, c.Bind(&request), nil)
	if err != nil {
		slog.ErrorContext(ctx, "Validation failed", slog.String(consts.Error, err.Error()))
		return c.JSON(response.Meta.Code, response)
	}

	data, err := h.healthService.CheckHealth(ctx, request.ToEntity())
	if err != nil {
		sc := utils.MapErrorToHTTP(err)
		response = dto.RestResponse(sc, nil, err)
		return c.JSON(sc, response)
	}

	response = dto.RestResponse(http.StatusOK, dto.ToCheckHealthResponse(data), nil)
	return c.JSON(response.Meta.Code, response)
}
