package health

import (
	"context"
	"fmt"
	"go-simple-template/internal/pkg/consts"
	"log/slog"
)

func (s *health) CheckHealth(ctx context.Context) error {
	const (
		funcName = "CheckHealth"
	)

	err := s.healthRepo.CheckHealth(ctx)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("Failed to CheckHealth", funcName), slog.String(consts.Error, err.Error()))
		return err
	}

	return nil
}
