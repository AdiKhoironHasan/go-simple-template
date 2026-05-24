package cache

import (
	"context"
	"time"
)

//go:generate mockgen -package mocks -source=token.go -destination=mocks/token_mock.go TokenBlacklist
type TokenBlacklist interface {
	Blacklist(ctx context.Context, token string, expiration time.Duration) error
	IsBlacklisted(ctx context.Context, token string) (bool, error)
}
