package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"go-simple-template/internal/adapter/inbound/rest/router"
	"go-simple-template/internal/pkg/config"
)

type server struct {
	srv *http.Server
}

func New(ctx context.Context, r *router.Dependencies) Server {
	rt := router.New(r)

	return &server{
		srv: &http.Server{
			Addr:    fmt.Sprintf(":%d", config.AppPort()),
			Handler: rt.Init(ctx),
		},
	}
}

func (s *server) Run(ctx context.Context) error {
	slog.InfoContext(ctx, "Server started", slog.String("address", s.srv.Addr))

	go func() {
		err := s.srv.ListenAndServe()
		if err != http.ErrServerClosed {
			slog.ErrorContext(ctx, "Failed to start the server", slog.String("error", err.Error()))
			return
		}
	}()

	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	slog.InfoContext(ctx, "Gracefully shutting down the rest server")

	err := s.srv.Shutdown(ctxShutDown)
	if err != nil && err != http.ErrServerClosed {
		slog.ErrorContext(ctx, "Failed to shutdown the server", slog.String("error", err.Error()))
		return err
	}

	return nil
}
