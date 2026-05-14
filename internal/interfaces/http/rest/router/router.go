package router

import (
	"context"
	_ "embed"

	"go-simple-template/internal/infrastructure"
	"go-simple-template/internal/pkg/config"

	healthRepo "go-simple-template/internal/adapter/repository/db/mongo/health"
	healthService "go-simple-template/internal/app/service/health"
	healthHandler "go-simple-template/internal/interfaces/http/rest/handler/health"
	"go-simple-template/internal/interfaces/http/rest/middleware"

	"github.com/labstack/echo/v4"
)

type router struct {
	factory *infrastructure.Factory
}

func New(opts ...Option) *router {
	r := &router{}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func (r *router) Init(ctx context.Context) *echo.Echo {
	e := echo.New()

	e.Debug = config.AppDebug()

	// add middleware, etc. here
	e.Use(
		middleware.MiddlewareCORS(),
		middleware.MiddlewareRecover(),
		middleware.MiddlewareLogger(),
		middleware.MiddlewareRequestID(),
	)

	// repository db
	healthDBRepo := healthRepo.New(r.factory.Mongodb)

	// service
	healthService := healthService.New(healthDBRepo)

	// handler
	healthHandler := healthHandler.New(healthService)

	// routes
	e.GET("/healthz", healthHandler.CheckHealth)

	return e
}
