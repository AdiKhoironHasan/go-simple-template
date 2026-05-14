package health

import (
	"net/http"

	"go-simple-template/internal/interfaces/http/rest/dto"

	"github.com/labstack/echo/v4"
)

func (h *health) CheckHealth(c echo.Context) error {
	var (
		ctx = c.Request().Context()
	)

	if err := h.healthService.CheckHealth(ctx); err != nil {
		response := dto.RestResponse(http.StatusInternalServerError, nil, nil)
		return c.JSON(response.Meta.Code, response)
	}

	response := dto.RestResponse(http.StatusOK, nil, nil)
	return c.JSON(response.Meta.Code, response)
}
