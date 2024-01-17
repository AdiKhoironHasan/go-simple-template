package service

import (
	"github.com/labstack/echo/v4"
)

func (s *service) Ping(ctx echo.Context) error {
	err := s.repo.Ping(ctx)
	if err != nil {
		logService.Error().Err(err).Msg("service ping failed")
		return err
	}

	return nil
}
