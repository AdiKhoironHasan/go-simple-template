package auth

import (
	"github.com/adikhoironhasan/go-simple-template/internal/core/port/outbound/cache"
	"github.com/adikhoironhasan/go-simple-template/internal/core/port/outbound/repository"
)

type auth struct {
	userRepo   repository.UserRepository
	tokenCache cache.TokenBlacklist
}

func New(
	userRepo repository.UserRepository,
	tokenCache cache.TokenBlacklist,
) *auth {
	return &auth{
		userRepo:   userRepo,
		tokenCache: tokenCache,
	}
}
