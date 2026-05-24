package auth

import (
	"log/slog"
	"net/http"

	"github.com/adikhoironhasan/go-simple-template/internal/adapter/inbound/rest/dto"
	"github.com/adikhoironhasan/go-simple-template/internal/adapter/inbound/rest/utils"
	"github.com/adikhoironhasan/go-simple-template/internal/pkg/consts"

	"github.com/labstack/echo/v4"
)

func (h *AuthHandler) Logout(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		request dto.AuthLogoutRequest
	)

	response, err := dto.ValidateRequest(ctx, c.Bind(&request), request.ValidateLogout())
	if err != nil {
		slog.ErrorContext(ctx, "Validation failed", slog.String(consts.Error, err.Error()))
		return c.JSON(response.Meta.Code, response)
	}

	err = h.authService.Logout(ctx, request.ToEntity())
	if err != nil {
		sc := utils.MapErrorToHTTP(err)
		response = dto.RestResponse(sc, nil, err)
		return c.JSON(sc, response)
	}

	response = dto.RestResponse(http.StatusOK, nil, nil)
	return c.JSON(response.Meta.Code, response)
}
