package server

import (
	"context"
	"fmt"
	"go-simple-template/config"
	"go-simple-template/internal/router"

	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type httpServer struct {
	srv *http.Server
}

func NewHttpServer(router *router.Router) Server {
	return &httpServer{
		srv: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", config.AppHost(), config.AppPort()),
			Handler: router.Init(),
		},
	}
}

func (s *httpServer) Run(ctx context.Context) error {
	log.Info().Msgf("Server started at http://%s:%d", config.AppHost(), config.AppPort())

	go func() {
		err := s.srv.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Error().Err(err).Msg("Failed to start the server")
		}
	}()

	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	log.Info().Msg("Server shutdown properly")

	err := s.srv.Shutdown(ctxShutDown)
	if err != nil && err != http.ErrServerClosed {
		log.Error().Err(err).Msg("Failed to shutdown the server")
		return err
	}

	return nil
}

func (s *httpServer) Done() {
	log.Info().Msg("Service http stopped")
}
