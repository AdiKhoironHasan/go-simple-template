package router

import (
	"go-simple-template/config"
	"go-simple-template/factory"
	"go-simple-template/internal/dto"
	"go-simple-template/internal/handler"
	"go-simple-template/internal/repository"
	"go-simple-template/internal/service"
	"net/http"

	tracemiddleware "go-simple-template/pkg/tracer/middleware"

	"github.com/labstack/echo/v4"
)

func New(opts ...Option) *Router {
	r := &Router{}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

// Router registers routes to be matched and dispatches a handler.
type Router struct {
	*factory.Factory
}

func (r *Router) Init() *echo.Echo {
	e := echo.New()

	e.Use(
		tracemiddleware.EchoMiddleware(config.AppName()),
	)

	// repository
	pingRepo := repository.NewPing(r.Db, r.Cache)

	// service
	pingService := service.NewPing(r.Logger, pingRepo, r.Storage)

	// handler
	pingHandler := handler.NewPing(pingService)

	// init routes
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, dto.ApiResponse{
			Code:    http.StatusOK,
			Message: "Welcome to Go Simple Template",
		})
	})

	// ping
	e.GET("/ping", pingHandler.Ping)

	return e
}
