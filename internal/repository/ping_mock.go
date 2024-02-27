package repository

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	Mock *mock.Mock
}

func NewRepositoryMock(mock *mock.Mock) RepositoryInterface {
	return &RepositoryMock{Mock: mock}
}

func (r *RepositoryMock) Ping(ctx context.Context) error {
	args := r.Mock.Called()
	return args.Error(0)
}
