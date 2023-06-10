package service

import (
	"go-simple-template/pkg/logger"
	"go-simple-template/repository"
)

type service struct {
	repo repository.RepositoryInterface
}

var (
	logService = logger.NewLogger().Logger.With().Str("pkg", "service").Logger()
)

func NewService(repo repository.RepositoryInterface) *service {
	return &service{
		repo: repo,
	}
}

type ServiceInterface interface {
	Ping() error
}
