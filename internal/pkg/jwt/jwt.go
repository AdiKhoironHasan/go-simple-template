package jwt

import (
	"errors"
	"log"
	"time"

	"go-simple-template/internal/core/domain/errs"
	"go-simple-template/internal/pkg/config"

	jwtgo "github.com/golang-jwt/jwt/v5"
)

func GenerateToken(payload UserCtx, isRefresh bool) (string, error) {
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

func ValidateToken(tokenString string, isRefreshToken bool) (*Claims, error) {
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
		log.Println("jwt err: ", err)
		return nil, errs.ErrInvalidToken
	}

	if !token.Valid {
		log.Println("jwt err: ", errs.ErrInvalidToken)
		return nil, errs.ErrInvalidToken
	}

	return claims, nil
}
