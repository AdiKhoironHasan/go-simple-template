package main

import (
	"fmt"
	"go-simple-template/config"
	"go-simple-template/database"
	"go-simple-template/handler"
	"go-simple-template/pkg/logger"
	"go-simple-template/repository"
	"go-simple-template/router"
	"go-simple-template/server"
	"go-simple-template/service"
	"os"
)

func main() {
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cfg := config.NewConfig(workDir)

	log := logger.NewLogger().Logger.With().Str("pkg", "main").Logger()

	db, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	router := router.NewRouter(handler)

	server := server.NewServer(cfg, router)

	log.Info().Msg(fmt.Sprintf("Server started at http://%s:%d", cfg.AppHost, cfg.AppPort))
	log.Fatal().Err(server.ListenAndServe()).Msg("Failed to start the server")
}
