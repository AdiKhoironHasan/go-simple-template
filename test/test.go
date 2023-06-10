package test

import (
	"go-simple-template/repository"

	"github.com/stretchr/testify/mock"
)

func CreateMock() (*mock.Mock, repository.RepositoryInterface) {
	m := mock.Mock{}
	repoMock := repository.NewRepositoryMock(&m)
	return &m, repoMock
}

func ResetMock(m *mock.Mock) {
	if m == nil {
		panic("mock is nil")
	}

	m.ExpectedCalls = nil
}
