package jwt

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"go-simple-template/internal/core/domain/entity"
	"go-simple-template/internal/pkg/config"
	errpkg "go-simple-template/internal/pkg/errs"

	jwtgo "github.com/golang-jwt/jwt/v5"
)

func GenerateToken(payload entity.UserCtx, isRefresh bool) (string, error) {
	var key []byte
	var expirationTime time.Time

	if isRefresh {
		key = []byte(config.AppRefreshKey())
		if len(key) == 0 {
			return "", errors.New("refresh key is not configured")
		}
		expirationTime = time.Now().Add(24 * time.Hour)
	} else {
		key = []byte(config.AppSecretKey())
		if len(key) == 0 {
			return "", errors.New("secret key is not configured")
		}
		expirationTime = time.Now().Add(2 * time.Hour)
	}

	claims := &Claims{
		UserCtx: payload,
		RegisteredClaims: jwtgo.RegisteredClaims{
			ExpiresAt: jwtgo.NewNumericDate(expirationTime),
		},
	}

	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func ValidateToken(ctx context.Context, tokenString string, isRefreshToken bool) (*Claims, error) {
	var (
		claims = &Claims{}
		key    = config.AppSecretKey()
	)

	if isRefreshToken {
		key = config.AppRefreshKey()
	}

	token, err := jwtgo.ParseWithClaims(tokenString, claims, func(token *jwtgo.Token) (any, error) {
		if _, ok := token.Method.(*jwtgo.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		if token.Method != jwtgo.SigningMethodHS256 {
			return nil, errors.New("unsupported signing algorithm")
		}
		return []byte(key), nil
	})
	if err != nil {
		slog.WarnContext(ctx, "JWT validation failed", slog.String("error", err.Error()))
		return nil, errpkg.NewUnauthorized(errpkg.ErrMsgInvalidToken)
	}

	if !token.Valid {
		slog.WarnContext(ctx, "JWT token is invalid")
		return nil, errpkg.NewUnauthorized(errpkg.ErrMsgInvalidToken)
	}

	return claims, nil
}
