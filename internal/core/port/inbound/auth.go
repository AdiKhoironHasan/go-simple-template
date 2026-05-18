package inbound

import (
	"context"

	"go-simple-template/internal/core/domain/entity"
)

type AuthService interface {
	Register(ctx context.Context, request entity.User) (*entity.User, error)
	Login(ctx context.Context, request entity.User) (*entity.AuthToken, error)
	RefreshToken(ctx context.Context, request entity.AuthToken) (*entity.AuthToken, error)
	Logout(ctx context.Context, request entity.AuthToken) error
	Profile(ctx context.Context, request entity.AuthToken) (*entity.User, error)
}
