package service

import (
	"context"
	"go-simple-template/pkg/tracer"
)

func (s *service) Ping(ctx context.Context) error {
	ctx, span := tracer.SpanStart(ctx, "Service.Ping")
	defer span.Finish()

	err := s.repo.Ping(ctx)
	if err != nil {
		span.AddError(err)
		logService.Error().Err(err).Msg("service ping failed")

		return err
	}

	return nil
}
