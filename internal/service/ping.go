package service

import (
	"context"
	"go-simple-template/internal/repository"
	"go-simple-template/pkg/storagex"
	"go-simple-template/pkg/tracer"

	"go.uber.org/zap"
)

type pingService struct {
	log     *zap.Logger
	repo    repository.PingRepository
	storage *storagex.Storage
}

func NewPing(log *zap.Logger, repo repository.PingRepository, Storage *storagex.Storage) *pingService {
	return &pingService{
		log:     log,
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
		s.log.Error("Failed to ping repository", zap.Error(err), zap.String("traceId", span.TraceId()))

		return err
	}

	return nil
}
