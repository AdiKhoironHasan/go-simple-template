package middleware

import (
	"net/http"
	"strings"

	"go-simple-template/internal/adapter/inbound/rest/dto"
	"go-simple-template/internal/pkg/jwt"

	ctxpkg "go-simple-template/internal/pkg/context"

	"github.com/labstack/echo/v4"
)

// MiddlewareJWT validates the JWT token in the Authorization header
func MiddlewareJWT() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				response := dto.RestResponse(http.StatusUnauthorized, nil, nil)
				return c.JSON(http.StatusUnauthorized, response)
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				response := dto.RestResponse(http.StatusUnauthorized, nil, nil)
				return c.JSON(http.StatusUnauthorized, response)
			}

			tokenString := parts[1]
			claims, err := jwt.ValidateToken(c.Request().Context(), tokenString, false)
			if err != nil {
				response := dto.RestResponse(http.StatusUnauthorized, nil, nil)
				return c.JSON(http.StatusUnauthorized, response)
			}

			// Add the user to context
			ctx := ctxpkg.SetUserCtx(c.Request().Context(), &claims.UserCtx)

			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
