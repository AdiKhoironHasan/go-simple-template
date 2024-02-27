package repository

import (
	"context"
	"go-simple-template/pkg/tracer"
)

func (r *repository) Ping(ctx context.Context) error {
	ctx, span := tracer.SpanStart(ctx, "Repository.Ping")
	defer span.Finish()

	span.AddEvents("pinging cache", tracer.SpanTagString("cache", "ping"))
	_, err := r.cache.Client.Ping(ctx)
	if err != nil {
		span.AddError(err)
		logRepo.Error().Err(err).Msg("cache ping failed")

		return err
	}

	if err := r.db.Exec("SELECT 1").Error; err != nil {
		span.AddError(err)
		logRepo.Error().Err(err).Msg("db ping failed")

		return err
	}

	return nil
}
