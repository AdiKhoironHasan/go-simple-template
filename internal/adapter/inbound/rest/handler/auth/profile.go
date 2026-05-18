package auth

import (
	"log/slog"
	"net/http"

	"go-simple-template/internal/adapter/inbound/rest/dto"
	"go-simple-template/internal/adapter/inbound/rest/utils"
	"go-simple-template/internal/core/domain/entity"

	ctxpkg "go-simple-template/internal/pkg/context"

	"github.com/labstack/echo/v4"
)

func (h *AuthHandler) Profile(c echo.Context) error {
	var (
		ctx = c.Request().Context()
	)

	userCtx, ok := ctxpkg.GetUserCtx(ctx)
	if !ok {
		slog.ErrorContext(ctx, "failed to get user context")
		response := dto.RestResponse(http.StatusUnauthorized, nil, nil)
		return c.JSON(response.Meta.Code, response)
	}

	data, err := h.authService.Profile(ctx, entity.AuthToken{Id: userCtx.Id})
	if err != nil {
		sc := utils.MapErrorToHTTP(err)
		response := dto.RestResponse(sc, nil, err)
		return c.JSON(sc, response)
	}

	resData := dto.ToProfileResponse(data)

	response := dto.RestResponse(http.StatusOK, resData, nil)
	return c.JSON(response.Meta.Code, response)
}
