package service

// import (
// 	"context"
// 	"go-simple-template/internal/repository"
// 	"go-simple-template/pkg/logger"
// 	"go-simple-template/pkg/storagex"

// 	"github.com/hibiken/asynq"
// )

// type pingService struct {
// 	repo    repository.PingRepository
// 	queue   *asynq.Client
// 	storage *storagex.Storage
// }

// var (
// 	logService = logger.NewLogger().Logger.With().Str("pkg", "service").Logger()
// )

// func NewPing() *pingService {
// 	return &pingService{}
// }

// type PingService interface {
// 	Ping(ctx context.Context) error
// }

// func (s *pingService) WithRepo(repo repository.PingRepository) *pingService {
// 	s.repo = repo
// 	return s
// }

// func (s *pingService) WithQueue(queue *asynq.Client) *pingService {
// 	s.queue = queue
// 	return s
// }

// func (s *pingService) WithStorage(Storage *storagex.Storage) *pingService {
// 	s.storage = Storage
// 	return s
// }
