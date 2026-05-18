package repository

import (
	"context"

	"go-simple-template/internal/core/domain/entity"
)

type UserRepository interface {
	Insert(ctx context.Context, request entity.User) (*entity.User, error)
	FindOne(ctx context.Context, request entity.User) (*entity.User, error)
	Update(ctx context.Context, request entity.User) (*entity.User, error)
}
