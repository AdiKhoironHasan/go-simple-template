package service

import (
	"go-simple-template/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	m        = mock.Mock{}
	repoMock = repository.NewRepositoryMock(&m)
)

func TestPing(t *testing.T) {
	service := NewService(repoMock)

	t.Run("error", func(t *testing.T) {
		m.On("Ping").Return(assert.AnError)
		err := service.Ping()

		assert.Error(t, err)
		assert.True(t, m.AssertCalled(t, "Ping"))

		// reset mock data
		m.ExpectedCalls = nil
		assert.Nil(t, m.ExpectedCalls)
	})

	t.Run("success", func(t *testing.T) {
		m.On("Ping").Return(nil)
		err := service.Ping()

		assert.NoError(t, err)
		assert.True(t, m.AssertCalled(t, "Ping"))

		// reset mock data
		m.ExpectedCalls = nil
		assert.Nil(t, m.ExpectedCalls)
	})
}
