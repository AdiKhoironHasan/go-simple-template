package auth

import (
	"context"
	"log/slog"
	"time"

	"go-simple-template/internal/core/domain/entity"
	"go-simple-template/internal/core/domain/errs"
	"go-simple-template/internal/pkg/consts"
	errpkg "go-simple-template/internal/pkg/errs"
	"go-simple-template/internal/pkg/jwt"
)

func (s *auth) RefreshToken(ctx context.Context, request entity.AuthToken) (*entity.AuthToken, error) {
	isBlacklisted, err := s.tokenCache.IsBlacklisted(ctx, request.RefreshToken)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to check blacklist", slog.String(consts.Error, err.Error()))
		return nil, errpkg.NewInternal(err, "failed to check blacklist")
	}

	if isBlacklisted {
		slog.ErrorContext(ctx, "Refresh token is blacklisted")
		return nil, errpkg.NewUnauthorized(errs.ErrMsgInvalidToken)
	}

	refreshToken, err := jwt.ValidateToken(request.RefreshToken, true)
	if err != nil {
		slog.ErrorContext(ctx, "Invalid refresh token", slog.String("error", err.Error()))
		return nil, errpkg.NewUnauthorized(errs.ErrMsgInvalidToken)
	}

	if refreshToken.ExpiresAt != nil {
		ttl := time.Until(refreshToken.ExpiresAt.Time)
		if ttl > 0 {
			if blErr := s.tokenCache.Blacklist(ctx, request.RefreshToken, ttl); blErr != nil {
				slog.ErrorContext(ctx, "Failed to blacklist old refresh token", slog.String(consts.Error, blErr.Error()))
			}
		}
	}

	jwtPayload := jwt.UserCtx{
		Id:    refreshToken.Id,
		Email: refreshToken.Email,
	}

	accessToken, newRefreshToken, err := s.generateAuthToken(ctx, jwtPayload)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to generate new tokens", slog.String("error", err.Error()))
		return nil, err
	}

	response := &entity.AuthToken{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}

	return response, nil
}
