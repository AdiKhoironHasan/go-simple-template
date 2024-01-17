package repository

import "github.com/labstack/echo/v4"

func (r *repository) Ping(ctx echo.Context) error {
	_, err := r.cache.Client.Ping()
	if err != nil {
		logRepo.Error().Err(err).Msg("cache ping failed")
		return err
	}

	if err := r.db.Exec("SELECT 1").Error; err != nil {
		logRepo.Error().Err(err).Msg("db ping failed")
		return err
	}

	return nil
}
