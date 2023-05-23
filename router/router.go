package router

import (
	"go-simple-template/dto"
	"go-simple-template/handler"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(h *handler.Handler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())

	InitRouter(e, h)

	return e
}

func InitRouter(e *echo.Echo, h *handler.Handler) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, dto.ApiResponse{
			Code:    http.StatusOK,
			Message: "Welcome to Go Simple Template",
		})
	})

	e.GET("/ping", h.Ping)
}
