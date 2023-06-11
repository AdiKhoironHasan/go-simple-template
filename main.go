package main

import (
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

	cache := cache.NewCache(redis)

	db, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	repo := repository.NewRepository().WithDB(db).WithCache(cache)
	service := service.NewService().WithRepo(repo)
	handler := handler.NewHandler().WithService(service)
	router := router.NewRouter(handler)

	server := server.NewServer(cfg, router)

	log.Info().Msgf("Server started at http://%s:%d", cfg.AppHost, cfg.AppPort)
	log.Fatal().Err(server.ListenAndServe()).Msg("Failed to start the server")
}
