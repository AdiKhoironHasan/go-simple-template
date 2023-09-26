package service

import (
	"go-simple-template/internal/repository"
	"go-simple-template/pkg/logger"
	"go-simple-template/pkg/storagex"

	"github.com/hibiken/asynq"
)

type service struct {
	repo    repository.RepositoryInterface
	queue   *asynq.Client
	storage *storagex.Storage
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

func (s *service) WithStorage(Storage *storagex.Storage) *service {
	s.storage = Storage
	return s
}
