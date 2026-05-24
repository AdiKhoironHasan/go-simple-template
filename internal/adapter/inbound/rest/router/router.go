package router

import (
	"context"

	"github.com/adikhoironhasan/go-simple-template/internal/adapter/inbound/rest/middleware"
	"github.com/adikhoironhasan/go-simple-template/internal/core/port/inbound"
	"github.com/adikhoironhasan/go-simple-template/internal/pkg/config"

	authHandler "github.com/adikhoironhasan/go-simple-template/internal/adapter/inbound/rest/handler/auth"
	healthHandler "github.com/adikhoironhasan/go-simple-template/internal/adapter/inbound/rest/handler/health"

	"github.com/labstack/echo/v4"
)

// Dependencies holds all pre-wired services that the router needs.
// The router does NOT do DI wiring — it only registers routes.
type Dependencies struct {
	HealthService inbound.HealthService
	AuthService   inbound.AuthService
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

	v1 := e.Group("/api/v1")

	// handler
	healthH := healthHandler.New(r.deps.HealthService)
	authH := authHandler.New(r.deps.AuthService)

	// routes
	e.GET("/healthz", healthH.CheckHealth)

	// auth routes
	v1.POST("/auth/register", authH.Register)
	v1.POST("/auth/login", authH.Login)
	v1.POST("/auth/refresh", authH.RefreshToken)
	v1.GET("/auth/profile", authH.Profile, middleware.MiddlewareJWT())
	v1.POST("/auth/logout", authH.Logout, middleware.MiddlewareJWT())

	return e
}
