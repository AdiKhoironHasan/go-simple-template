package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

func GinMiddleware(serviceName string) gin.HandlerFunc {
	return otelgin.Middleware(serviceName)
}

func EchoMiddleware(serviceName string) echo.MiddlewareFunc {
	return otelecho.Middleware(serviceName)
}
