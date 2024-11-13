package router

import (
	"fmt"
	"go-simple-template/config"
	"go-simple-template/factory"
	"go-simple-template/internal/handler"
	"go-simple-template/internal/repository"
	"go-simple-template/internal/service"

	"go-simple-template/pkg/api/rest"
	tracemiddleware "go-simple-template/pkg/tracer/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
		middleware.Recover(),
		tracemiddleware.EchoMiddleware(config.AppName()),
	)

	// repository
	pingRepo := repository.NewPing(r.Db, r.Cache)

	// service
	pingService := service.NewPing(r.Logger, pingRepo, r.Storage, &r.RabbitMQ)

	// handler
	pingHandler := handler.NewPing(pingService)

	// init routes
	e.GET("/", func(c echo.Context) error {
		response := rest.ApiResponse().WithMessage(fmt.Sprintf("Welcome to %s", config.AppName()))
		return c.JSON(response.Meta.Code, response)
	})

	r.pingRouter(e, pingHandler)

	return e
}
