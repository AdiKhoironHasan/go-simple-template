package token

import (
	"context"
	"log/slog"
	"time"

	"go-simple-template/internal/pkg/consts"
)

func (c *cacheRepo) Blacklist(ctx context.Context, token string, expiration time.Duration) error {
	if expiration <= 0 {
		expiration = time.Second
	}
	key := blacklistPrefix + hashToken(token)
	err := c.client.Set(ctx, key, "1", expiration).Err()
	if err != nil {
		slog.ErrorContext(ctx, "Failed to blacklist token", slog.String(consts.Error, err.Error()))
		return err
	}
	return nil
}
