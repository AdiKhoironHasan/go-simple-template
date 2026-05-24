package jwt

import (
	"go-simple-template/internal/core/domain/entity"

	jwtgo "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserCtx entity.UserCtx
	jwtgo.RegisteredClaims
}

// GetExpirationTime implements the Claims interface.
func (c *Claims) GetExpirationTime() (*jwtgo.NumericDate, error) {
	return c.ExpiresAt, nil
}

// GetNotBefore implements the Claims interface.
func (c *Claims) GetNotBefore() (*jwtgo.NumericDate, error) {
	return c.NotBefore, nil
}

// GetIssuedAt implements the Claims interface.
func (c *Claims) GetIssuedAt() (*jwtgo.NumericDate, error) {
	return c.IssuedAt, nil
}

// GetAudience implements the Claims interface.
func (c *Claims) GetAudience() (jwtgo.ClaimStrings, error) {
	return c.Audience, nil
}

// GetIssuer implements the Claims interface.
func (c *Claims) GetIssuer() (string, error) {
	return c.Issuer, nil
}

// GetSubject implements the Claims interface.
func (c *Claims) GetSubject() (string, error) {
	return c.Subject, nil
}
