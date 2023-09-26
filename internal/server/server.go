package server

import (
	"fmt"
	"go-simple-template/config"
	"net/http"
)

func NewServer(cfg *config.Config, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.AppHost, cfg.AppPort),
		Handler: handler,
	}
}
