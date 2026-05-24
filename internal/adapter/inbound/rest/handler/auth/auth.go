package auth

import (
	"github.com/adikhoironhasan/go-simple-template/internal/core/port/inbound"
)

type AuthHandler struct {
	authService inbound.AuthService
}

func New(authService inbound.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}
