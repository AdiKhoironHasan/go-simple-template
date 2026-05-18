package auth

import (
	"go-simple-template/internal/core/port/inbound"
)

type AuthHandler struct {
	authService inbound.AuthService
}

func New(authService inbound.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}


