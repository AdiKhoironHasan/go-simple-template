package middleware

import (
	"context"
	"log/slog"
	"os"

	"go-simple-template/internal/pkg/consts"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MiddlewareLogger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true,
		Skipper:     skipper,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

			attrs := []slog.Attr{
				slog.String("uri", v.URI),
				slog.Int("status", v.Status),
				slog.String("method", v.Method),
				slog.String("latency", v.Latency.String()),
				slog.String("request_id", c.Response().Header().Get(echo.HeaderXRequestID)),
			}

			if v.Error != nil {
				attrs = append(attrs, slog.String(consts.Error, v.Error.Error()))
			}

			logger.LogAttrs(
				c.Request().Context(),
				slog.LevelInfo,
				"REQUEST",
				attrs...,
			)

			return nil
		},
	})
}

func skipper(c echo.Context) bool {
	// Implement your logic to skip logging for certain requests
	// For example, skip logging for health check endpoints
	return c.Request().URL.Path == "/healthz"
}

func MiddlewareCORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"*",
		},
		AllowMethods: []string{
			echo.GET,
			echo.POST,
			echo.PUT,
			echo.PATCH,
			echo.DELETE,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			"Authorization",
		},
		ExposeHeaders: []string{
			echo.HeaderXRequestID,
		},
	})
}

func MiddlewareRecover() echo.MiddlewareFunc {
	return middleware.RecoverWithConfig(middleware.DefaultRecoverConfig)
}

func MiddlewareRequestID() echo.MiddlewareFunc {
	config := middleware.DefaultRequestIDConfig

	// Set the request ID in the request context
	// This allows other parts of the application to access the request ID in the request context
	config.RequestIDHandler = func(c echo.Context, requestId string) {
		ctx := context.WithValue(c.Request().Context(), consts.CtxRequestId, requestId)
		c.SetRequest(c.Request().WithContext(ctx))
	}

	return middleware.RequestIDWithConfig(config)
}
