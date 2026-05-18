package auth

import (
	"log/slog"
	"net/http"

	"go-simple-template/internal/adapter/inbound/rest/dto"
	"go-simple-template/internal/adapter/inbound/rest/utils"
	"go-simple-template/internal/pkg/consts"

	"github.com/labstack/echo/v4"
)

func (h *AuthHandler) RefreshToken(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		request dto.AuthRefreshRequest
	)

	response, err := dto.ValidateRequest(ctx, c.Bind(&request), request.ValidateRefreshToken())
	if err != nil {
		slog.ErrorContext(ctx, "Validation failed", slog.String(consts.Error, err.Error()))
		return c.JSON(response.Meta.Code, response)
	}

	data, err := h.authService.RefreshToken(ctx, request.ToEntity())
	if err != nil {
		sc := utils.MapErrorToHTTP(err)
		response = dto.RestResponse(sc, nil, err)
		return c.JSON(sc, response)
	}

	resData := dto.ToAuthTokenResponse(data)

	response = dto.RestResponse(http.StatusOK, resData, nil)
	return c.JSON(response.Meta.Code, response)
}
