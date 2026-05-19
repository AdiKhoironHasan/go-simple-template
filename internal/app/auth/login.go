package auth

import (
	"context"
	"log/slog"

	"go-simple-template/internal/core/domain/entity"
	"go-simple-template/internal/core/domain/errs"
	"go-simple-template/internal/pkg/crypto"
	errpkg "go-simple-template/internal/pkg/errs"
	"go-simple-template/internal/pkg/jwt"
)

func (s *auth) Login(ctx context.Context, request entity.User) (*entity.AuthToken, error) {
	user, err := s.userRepo.FindOne(ctx, request)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to find user", slog.String("error", err.Error()))

		if errpkg.GetCode(err) == errpkg.ErrNotFound {
			return nil, errpkg.NewUnauthorized(errs.ErrMsgInvalidCredentials)
		}

		return nil, err
	}

	// check if password is correct
	if !crypto.CheckPasswordHash(request.Password, user.Password) {
		slog.ErrorContext(ctx, "Incorrect password", slog.String("user_id", user.Id))
		return nil, errpkg.NewUnauthorized(errs.ErrMsgInvalidCredentials)
	}

	jwtPayload := jwt.UserCtx{
		Id:    user.Id,
		Email: user.Email,
	}

	// generate access token
	accessToken, refreshToken, err := s.generateAuthToken(ctx, jwtPayload)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to generate tokens", slog.String("error", err.Error()))
		return nil, err
	}

	response := &entity.AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func (s *auth) generateAuthToken(ctx context.Context, payload jwt.UserCtx) (string, string, error) {
	accessToken, err := jwt.GenerateToken(payload, false)
	if err != nil {
		slog.ErrorContext(ctx, "error generating access token", slog.String("error", err.Error()))
		return "", "", err
	}

	refreshToken, err := jwt.GenerateToken(payload, true)
	if err != nil {
		slog.ErrorContext(ctx, "error generating refresh token", slog.String("error", err.Error()))
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
