package service

func (s *service) Ping() error {
	return s.repo.Ping()
}
