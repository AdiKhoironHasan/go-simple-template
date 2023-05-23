package service

import "go-simple/repository"

type service struct {
	repo repository.RepositoryInterface
}

func NewService(repo repository.RepositoryInterface) *service {
	return &service{
		repo: repo,
	}
}

type ServiceInterface interface {
	Ping() error
}
