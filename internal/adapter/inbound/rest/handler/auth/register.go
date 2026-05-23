package auth

import (
	"log/slog"
	"net/http"

	"go-simple-template/internal/adapter/inbound/rest/dto"
	"go-simple-template/internal/adapter/inbound/rest/utils"
	"go-simple-template/internal/pkg/consts"

	"github.com/labstack/echo/v4"
)

func (h *AuthHandler) Register(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		request dto.AuthRegisterRequest
	)

	response, err := dto.ValidateRequest(ctx, c.Bind(&request), request.ValidateRegister())
	if err != nil {
		slog.ErrorContext(ctx, "Validation failed", slog.String(consts.Error, err.Error()))
		return c.JSON(response.Meta.Code, response)
	}

	userEntity := request.ToEntity()

	data, err := h.authService.Register(ctx, userEntity)
	if err != nil {
		sc := utils.MapErrorToHTTP(err)
		response = dto.RestResponse(sc, nil, err)
		return c.JSON(sc, response)
	}

	resData := dto.ToRegisterResponse(data)

	response = dto.RestResponse(http.StatusCreated, resData, nil)
	return c.JSON(response.Meta.Code, response)
}
