package main

import (
	"fmt"
	"go-simple-template/config"
	"go-simple-template/database"
	"go-simple-template/handler"
	"go-simple-template/repository"
	"go-simple-template/router"
	"go-simple-template/service"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	if errEnv := godotenv.Load(); errEnv != nil {
		log.Fatal().Err(errEnv).Msg("Failed to load .env file")
	}
	cfg := config.NewConfig()

	dbConn, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	repo := repository.NewRepository(dbConn)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	router := router.NewRouter(handler)

	server := NewServer(cfg, router)

	log.Info().Msg(fmt.Sprintf("Server started at  http://%s:%d", cfg.AppHost, cfg.AppPort))
	log.Fatal().Err(server.ListenAndServe()).Msg("Failed to start the server")
}
