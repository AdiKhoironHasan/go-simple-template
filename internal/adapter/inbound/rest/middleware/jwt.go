package middleware

import (
	"net/http"
	"strings"

	"go-simple-template/internal/core/domain/entity"
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
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header format")
			}

			tokenString := parts[1]
			claims, err := jwt.ValidateToken(tokenString, false)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
			}

			// Add the user to context
			ctx := ctxpkg.SetUserCtx(c.Request().Context(), &entity.UserCtx{
				Id:    claims.UserCtx.Id,
				Email: claims.UserCtx.Email,
			})

			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
