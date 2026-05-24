package auth

import (
	"context"
	"log/slog"
	"time"

	"go-simple-template/internal/core/domain/entity"
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
		return nil, errpkg.NewUnauthorized(errpkg.ErrMsgInvalidToken)
	}

	claims, err := jwt.ValidateToken(ctx, request.RefreshToken, true)
	if err != nil {
		slog.ErrorContext(ctx, "Invalid refresh token", slog.String(consts.Error, err.Error()))
		return nil, errpkg.NewUnauthorized(errpkg.ErrMsgInvalidToken)
	}

	if claims.ExpiresAt != nil {
		ttl := time.Until(claims.ExpiresAt.Time)
		if ttl > 0 {
			if blErr := s.tokenCache.Blacklist(ctx, request.RefreshToken, ttl); blErr != nil {
				slog.ErrorContext(ctx, "Failed to blacklist old refresh token", slog.String(consts.Error, blErr.Error()))
			}
		}
	}

	payload := entity.UserCtx{Id: claims.UserCtx.Id, Email: claims.UserCtx.Email}

	accessToken, err := jwt.GenerateToken(payload, false)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to generate access token", slog.String(consts.Error, err.Error()))
		return nil, err
	}

	newRefreshToken, err := jwt.GenerateToken(payload, true)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to generate refresh token", slog.String(consts.Error, err.Error()))
		return nil, err
	}

	return &entity.AuthToken{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
