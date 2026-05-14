package health

import (
	"context"
	"go-simple-template/internal/pkg/consts"
	"log/slog"
)

func (db *repository) CheckHealth(ctx context.Context) error {
	// ping to mongoDB
	err := db.client.Ping(ctx, nil)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to ping MongoDB", slog.String(consts.Error, err.Error()))
		return err
	}

	return nil
}
