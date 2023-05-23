package router

import (
	"go-simple/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewRouter is a constructor will initialize Router.
func NewRouter(h *handler.Handler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello World")
	})
	e.GET("/ping", h.Ping)

	return e
}
