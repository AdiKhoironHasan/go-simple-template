package auth

import (
	"context"
	"log/slog"
	"time"

	"github.com/adikhoironhasan/go-simple-template/internal/core/domain/entity"
	"github.com/adikhoironhasan/go-simple-template/internal/pkg/consts"
	errpkg "github.com/adikhoironhasan/go-simple-template/internal/pkg/errs"
	"github.com/adikhoironhasan/go-simple-template/internal/pkg/jwt"
)

func (s *auth) Logout(ctx context.Context, request entity.AuthToken) error {
	claims, err := jwt.ValidateToken(ctx, request.RefreshToken, true)
	if err != nil {
		slog.ErrorContext(ctx, "Invalid refresh token", slog.String(consts.Error, err.Error()))
		return errpkg.NewUnauthorized(errpkg.ErrMsgInvalidToken)
	}

	if claims.ExpiresAt == nil {
		slog.ErrorContext(ctx, "Refresh token missing expiration")
		return errpkg.NewUnauthorized(errpkg.ErrMsgInvalidToken)
	}

	ttl := time.Until(claims.ExpiresAt.Time)
	if ttl <= 0 {
		ttl = time.Second
	}

	err = s.tokenCache.Blacklist(ctx, request.RefreshToken, ttl)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to blacklist token", slog.String(consts.Error, err.Error()))
		return err
	}

	slog.InfoContext(ctx, "User logged out successfully")

	return nil
}
