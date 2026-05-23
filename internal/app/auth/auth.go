package auth

import (
	"go-simple-template/internal/core/port/inbound"
	"go-simple-template/internal/core/port/outbound/cache"
	"go-simple-template/internal/core/port/outbound/repository"
)

type auth struct {
	userRepo   repository.UserRepository
	tokenCache cache.Token
}

func New(
	userRepo repository.UserRepository,
	tokenCache cache.Token,
) inbound.AuthService {
	return &auth{
		userRepo:   userRepo,
		tokenCache: tokenCache,
	}
}
