package router

import (
	"go-simple-template/factory"
	"go-simple-template/internal/dto"
	"go-simple-template/internal/handler"
	"go-simple-template/internal/repository"
	"go-simple-template/internal/service"
	"net/http"

	tracemiddleware "go-simple-template/pkg/tracer/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(opts ...Option) *Router {
	r := &Router{
		// Router: echo.New().Router(),
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

// Router registers routes to be matched and dispatches a handler.
type Router struct {
	// *echo.Router
	*factory.Factory
}

func (r *Router) Init() *echo.Echo {
	e := echo.New()

	e.Use(
		middleware.Logger(),
		tracemiddleware.EchoMiddleware("svcname"),
	)

	// repository
	pingRepo := repository.NewPing().WithDB(r.Db).WithCache(r.Cache)

	// service
	pingService := service.NewPing(pingRepo, r.Storage)

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
