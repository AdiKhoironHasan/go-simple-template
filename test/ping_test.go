package test

import (
	"encoding/json"
	"fmt"
	"go-simple-template/config"
	"go-simple-template/internal/dto"
	"go-simple-template/internal/handler"
	"go-simple-template/internal/router"
	"go-simple-template/internal/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	m, repoMock := CreateMock()

	err := godotenv.Load("../.env")
	assert.NoError(t, err)

	cfg := config.NewConfig()

	service := service.NewService().WithRepo(repoMock)
	handler := handler.NewHandler().WithService(service)
	router := router.NewRouter(handler, cfg)

	t.Run("success", func(t *testing.T) {
		m.On("Ping").Return(nil)

		requestBody := strings.NewReader(`{}`)
		request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://%s:%d/ping", cfg.AppHost, cfg.AppPort), requestBody)
		request.Header.Add("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)

		response := recorder.Result()
		body, err := io.ReadAll(response.Body)
		assert.NoError(t, err)

		var responseBody dto.ApiResponse
		json.Unmarshal(body, &responseBody)

		ResetMock(m)
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, "pong", responseBody.Message)
	})

	t.Run("error", func(t *testing.T) {
		m.On("Ping").Return(assert.AnError)

		requestBody := strings.NewReader(`{}`)
		request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://%s:%d/ping", cfg.AppHost, cfg.AppPort), requestBody)
		request.Header.Add("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)

		response := recorder.Result()
		body, err := io.ReadAll(response.Body)
		assert.NoError(t, err)

		var responseBody dto.ApiResponse
		json.Unmarshal(body, &responseBody)

		ResetMock(m)
		assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
		assert.Equal(t, assert.AnError.Error(), responseBody.Message)
	})
}
