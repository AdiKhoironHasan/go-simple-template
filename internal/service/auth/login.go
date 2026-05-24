package auth

import (
	"context"
	"log/slog"

	"go-simple-template/internal/core/domain/entity"
	"go-simple-template/internal/pkg/crypto"
	errpkg "go-simple-template/internal/pkg/errs"
	"go-simple-template/internal/pkg/jwt"
)

func (s *auth) Login(ctx context.Context, request entity.User) (*entity.AuthToken, error) {
	user, err := s.userRepo.FindOne(ctx, request)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to find user", slog.String("error", err.Error()))

		if errpkg.GetCode(err) == errpkg.ErrNotFound {
			return nil, errpkg.NewUnauthorized(errpkg.ErrMsgInvalidCredentials)
		}

		return nil, err
	}

	if !crypto.CheckPasswordHash(request.Password, user.Password) {
		slog.ErrorContext(ctx, "Incorrect password", slog.String("user_id", user.Id))
		return nil, errpkg.NewUnauthorized(errpkg.ErrMsgInvalidCredentials)
	}

	payload := entity.UserCtx{Id: user.Id, Email: user.Email}

	accessToken, err := jwt.GenerateToken(payload, false)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to generate access token", slog.String("error", err.Error()))
		return nil, err
	}

	refreshToken, err := jwt.GenerateToken(payload, true)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to generate refresh token", slog.String("error", err.Error()))
		return nil, err
	}

	return &entity.AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
