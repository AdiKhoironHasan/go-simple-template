package service

import (
	"go-simple-template/pkg/logger"
	"go-simple-template/repository"

	"github.com/hibiken/asynq"
)

type service struct {
	repo  repository.RepositoryInterface
	queue *asynq.Client
}

var (
	logService = logger.NewLogger().Logger.With().Str("pkg", "service").Logger()
)

func NewService() *service {
	return &service{}
}

type ServiceInterface interface {
	Ping() error
}

func (s *service) WithRepo(repo repository.RepositoryInterface) *service {
	s.repo = repo
	return s
}

func (s *service) WithQueue(queue *asynq.Client) *service {
	s.queue = queue
	return s
}
