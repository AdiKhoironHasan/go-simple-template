package service

import (
	"context"
	"go-simple-template/internal/repository"
	"go-simple-template/pkg/logger"
	"go-simple-template/pkg/storagex"
	"go-simple-template/pkg/tracer"
)

type pingService struct {
	repo    repository.PingRepository
	storage *storagex.Storage
}

var (
	logService = logger.NewLogger().Logger.With().Str("pkg", "service").Logger()
)

func NewPing(repo repository.PingRepository, Storage *storagex.Storage) *pingService {
	return &pingService{
		repo:    repo,
		storage: Storage,
	}
}

type PingService interface {
	Ping(ctx context.Context) error
}

func (s *pingService) Ping(ctx context.Context) error {
	ctx, span := tracer.SpanStart(ctx, "Service.Ping")
	defer span.Finish()

	err := s.repo.Ping(ctx)
	if err != nil {
		span.AddError(err)
		logService.Error().Err(err).Str("traceId", span.TraceId()).Msg("service ping failed")

		return err
	}

	return nil
}
