package service

import (
	"go-simple-template/test"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	m, repoMock := test.CreateMock()

	service := NewService().WithRepo(repoMock)

	t.Run("error", func(t *testing.T) {
		echoCtx := echo.New().NewContext(nil, nil)

		m.On("Ping").Return(assert.AnError)
		err := service.Ping(echoCtx)

		assert.Error(t, err)
		assert.True(t, m.AssertCalled(t, "Ping"))

		test.ResetMock(m)
		assert.Nil(t, m.ExpectedCalls)
	})

	t.Run("success", func(t *testing.T) {
		echoCtx := echo.New().NewContext(nil, nil)

		m.On("Ping").Return(nil)
		err := service.Ping(echoCtx)

		assert.NoError(t, err)
		assert.True(t, m.AssertCalled(t, "Ping"))

		test.ResetMock(m)
		assert.Nil(t, m.ExpectedCalls)
	})
}
