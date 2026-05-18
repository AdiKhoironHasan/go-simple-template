package jwt

import (
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
		expirationTime = time.Now().Add(24 * time.Hour)
	} else {
		key = []byte(config.AppSecretKey())
		expirationTime = time.Now().Add(2 * time.Hour)
	}

	claims := &Claims{
		UserCtx:   payload,
		RevokedAt: jwtgo.NewNumericDate(time.Now().AddDate(0, 0, 7)),
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

	if claims.RevokedAt != nil && claims.RevokedAt.Time.Before(time.Now()) {
		log.Println("jwt err: token revoked at ", claims.RevokedAt.Time.Format(time.DateTime))
		return nil, errs.ErrInvalidToken
	}

	return claims, nil
}
