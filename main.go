package main

import (
	"fmt"
	"go-simple-template/cache"
	"go-simple-template/cache/redis"
	"go-simple-template/config"
	"go-simple-template/database"
	"go-simple-template/handler"
	"go-simple-template/pkg/logger"
	"go-simple-template/repository"
	"go-simple-template/router"
	"go-simple-template/server"
	"go-simple-template/service"

	"github.com/joho/godotenv"
)

func main() {
	log := logger.NewLogger().Logger.With().Str("pkg", "main").Logger()

	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("Failed to load env file")
		panic(err)
	}

	cfg := config.NewConfig()

	redis := redis.NewRedis(cfg)
	log.Info().Msg("Redis connected")

	cache := cache.NewCache(redis)
	log.Info().Msg("Cache connected")

	db, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	repo := repository.NewRepository(db, cache)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	router := router.NewRouter(handler)

	server := server.NewServer(cfg, router)

	log.Info().Msg(fmt.Sprintf("Server started at http://%s:%d", cfg.AppHost, cfg.AppPort))
	log.Fatal().Err(server.ListenAndServe()).Msg("Failed to start the server")
}
