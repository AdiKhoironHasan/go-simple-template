package auth

import (
	"context"
	"log/slog"

	"go-simple-template/internal/core/domain/entity"
	"go-simple-template/internal/pkg/crypto"
	errpkg "go-simple-template/internal/pkg/errs"
)

func (s *auth) Register(ctx context.Context, request entity.User) (*entity.User, error) {
	user, err := s.userRepo.FindOne(ctx, request)
	if err != nil && errpkg.GetCode(err) != errpkg.ErrNotFound {
		slog.ErrorContext(ctx, "Failed to find user", slog.String("error", err.Error()))
		return nil, err
	}

	if user != nil {
		slog.ErrorContext(ctx, "User already exists", slog.String("email", request.Email))
		return nil, errpkg.NewConflict(errpkg.ErrMsgEmailAlreadyExists)
	}

	hashedPassword, err := crypto.HashPassword(request.Password)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to hash password", slog.String("error", err.Error()))
		return nil, err
	}

	request.Password = hashedPassword

	response, err := s.userRepo.Insert(ctx, request)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to create user", slog.String("error", err.Error()))
		return nil, err
	}

	return response, nil
}
