package service

import (
	"context"
	"go-simple-template/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	ctx := context.Background()

	m, repoMock := test.CreateMock()

	service := NewService().WithRepo(repoMock)

	t.Run("error", func(t *testing.T) {
		m.On("Ping").Return(assert.AnError)
		err := service.Ping(ctx)

		assert.Error(t, err)
		assert.True(t, m.AssertCalled(t, "Ping"))

		test.ResetMock(m)
		assert.Nil(t, m.ExpectedCalls)
	})

	t.Run("success", func(t *testing.T) {
		m.On("Ping").Return(nil)
		err := service.Ping(ctx)

		assert.NoError(t, err)
		assert.True(t, m.AssertCalled(t, "Ping"))

		test.ResetMock(m)
		assert.Nil(t, m.ExpectedCalls)
	})
}
