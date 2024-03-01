package router

import (
	"go-simple-template/config"
	"go-simple-template/internal/dto"
	"go-simple-template/internal/handler"
	"net/http"

	tracemiddleware "go-simple-template/pkg/tracer/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(h *handler.Handler, cfg *config.Config) *echo.Echo {
	e := echo.New()
	e.Use(
		middleware.Logger(),
		tracemiddleware.EchoMiddleware(cfg.AppName),
	)

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
