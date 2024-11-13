package router

import (
	"go-simple-template/internal/handler"

	"github.com/labstack/echo/v4"
)

func (r *Router) pingRouter(e *echo.Echo, h *handler.PingHandler) {
	e.GET("/ping", h.Ping)
}
