package test

import (
	"encoding/json"
	"fmt"
	"go-simple-template/config"
	"go-simple-template/dto"
	"go-simple-template/handler"
	"go-simple-template/router"
	"go-simple-template/service"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	m, repoMock := CreateMock()

	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// mundur 1 direktori ke belakang as ../
	rootDir := filepath.Dir(workDir)

	cfg := config.NewConfig(rootDir)

	service := service.NewService(repoMock)
	handler := handler.NewHandler(service)
	router := router.NewRouter(handler)

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
