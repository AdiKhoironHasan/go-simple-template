package router

import (
	"context"

	"github.com/labstack/echo/v4"
)

type Router interface {
	Init(ctx context.Context) *echo.Echo
}
