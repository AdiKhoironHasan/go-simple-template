package health

import (
	"context"
	"github.com/adikhoironhasan/go-simple-template/internal/pkg/consts"
	"log/slog"
)

func (c *cacheRepo) CheckHealth(ctx context.Context) error {
	err := c.client.Ping(ctx).Err()
	if err != nil {
		slog.ErrorContext(ctx, "Failed to ping Redis", slog.String(consts.Error, err.Error()))
		return err
	}

	return nil
}
