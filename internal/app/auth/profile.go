package auth

import (
	"context"
	"log/slog"

	"go-simple-template/internal/core/domain/entity"
)

func (s *auth) Profile(ctx context.Context, request entity.AuthToken) (*entity.User, error) {
	profile, err := s.userRepo.FindOne(ctx, entity.User{Base: entity.Base{Id: request.Id}})
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get user profile", slog.String("error", err.Error()))
		return nil, err
	}

	return profile, nil
}
