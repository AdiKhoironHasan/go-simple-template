package repository

import (
	"context"

	"github.com/adikhoironhasan/go-simple-template/internal/core/domain/entity"
)

//go:generate mockgen -package mocks -source=user.go -destination=mocks/user_mock.go UserRepository
type UserRepository interface {
	Insert(ctx context.Context, request entity.User) (*entity.User, error)
	FindOne(ctx context.Context, request entity.User) (*entity.User, error)
}
