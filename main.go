package main

import (
	"fmt"
	"go-simple/config"
	"go-simple/database"
	"go-simple/handler"
	"go-simple/repository"
	"go-simple/router"
	"go-simple/service"
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config.Config{
		AppConfig: config.AppConfig{
			AppHost: "localhost",
			AppPort: 8000,
		},
		DBconfig: config.DBconfig{
			DBdriver: "mysql",
			DBhost:   "localhost",
			DBport:   "3306",
			DBuser:   "root",
			DBname:   "go-simple",
		},
	}

	dbConn, err := database.NewConnection(&cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	// dependency injection
	repo := repository.NewRepository(dbConn)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	router := router.NewRouter(handler)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.AppHost, cfg.AppPort),
		Handler: router,
	}

	log.Info().Msg(fmt.Sprintf("Server started at  http://%s:%d", cfg.AppHost, cfg.AppPort))
	log.Fatal().Err(server.ListenAndServe()).Msg("Failed to start the server")
}
