package service

func (s *service) Ping() error {
	err := s.repo.Ping()
	if err != nil {
		logService.Error().Err(err).Msg("service ping failed")
		return err
	}

	return nil
}
