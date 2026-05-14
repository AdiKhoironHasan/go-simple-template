package router

import (
	"context"

	"go-simple-template/internal/adapter/inbound/rest/middleware"
	"go-simple-template/internal/pkg/config"
	"go-simple-template/internal/core/port/inbound"

	healthHandler "go-simple-template/internal/adapter/inbound/rest/handler/health"

	"github.com/labstack/echo/v4"
)

// Dependencies holds all pre-wired services that the router needs.
// The router does NOT do DI wiring — it only registers routes.
type Dependencies struct {
	HealthService inbound.HealthService
}

type router struct {
	deps *Dependencies
}

func New(deps *Dependencies) *router {
	return &router{deps: deps}
}

func (r *router) Init(ctx context.Context) *echo.Echo {
	e := echo.New()

	e.Debug = config.AppDebug()

	// add middleware
	e.Use(
		middleware.MiddlewareCORS(),
		middleware.MiddlewareRecover(),
		middleware.MiddlewareLogger(),
		middleware.MiddlewareRequestID(),
	)

	// handler
	healthH := healthHandler.New(r.deps.HealthService)

	// routes
	e.GET("/healthz", healthH.CheckHealth)

	return e
}
