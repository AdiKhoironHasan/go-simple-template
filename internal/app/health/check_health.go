package health

import (
	"context"
	"fmt"
	"go-simple-template/internal/core/domain/entity"
	"go-simple-template/internal/pkg/consts"
	"log/slog"

	"golang.org/x/sync/errgroup"
)

func (s *service) CheckHealth(ctx context.Context, req entity.CheckHealth) (*entity.CheckHealth, error) {
	const (
		funcName = "CheckHealth"
	)

	response := entity.CheckHealth{}

	eg, egCtx := errgroup.WithContext(ctx)

	if req.MongoDB {
		eg.Go(func() error {
			err := s.healthRepo.CheckHealth(egCtx)
			if err != nil {
				slog.ErrorContext(egCtx, fmt.Sprintf("Failed to %s MongoDB", funcName), slog.String(consts.Error, err.Error()))
				return err
			}

			response.MongoDB = true
			return nil
		})
	}

	if req.Redis {
		eg.Go(func() error {
			err := s.healthCache.CheckHealth(egCtx)
			if err != nil {
				slog.ErrorContext(egCtx, fmt.Sprintf("Failed to %s Redis", funcName), slog.String(consts.Error, err.Error()))
				return err
			}

			response.Redis = true
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return &response, nil
}
