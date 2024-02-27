package main

import (
	"go-simple-template/config"
	"go-simple-template/internal/database"
	"go-simple-template/internal/handler"
	"go-simple-template/internal/repository"
	"go-simple-template/internal/router"
	"go-simple-template/internal/server"
	"go-simple-template/internal/service"
	"go-simple-template/pkg/cachex"
	"go-simple-template/pkg/cachex/redis"
	"go-simple-template/pkg/logger"
	"go-simple-template/pkg/storagex"
	"go-simple-template/pkg/storagex/minio"
	"go-simple-template/pkg/tracer"

	"github.com/joho/godotenv"
)

func main() {
	log := logger.NewLogger().Logger.With().Str("pkg", "main").Logger()

	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("Failed to load env file")
		panic(err)
	}

	cfg := config.NewConfig()

	tp, err := tracer.JaegerTraceProvider(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create trace provider")
	}

	tracer.RegisterTracer(tp)

	redis := redis.NewRedis(cfg)

	cache := cachex.NewCache(redis)

	db, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	minio, err := minio.NewMinio(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to minio")
	}
	storage := storagex.NewStorage(minio)

	repo := repository.NewRepository().WithDB(db).WithCache(cache)
	service := service.NewService().WithRepo(repo).WithStorage(storage)
	handler := handler.NewHandler().WithService(service)
	router := router.NewRouter(handler)

	server := server.NewServer(cfg, router)

	log.Info().Msgf("Server started at http://%s:%d", cfg.AppHost, cfg.AppPort)
	log.Fatal().Err(server.ListenAndServe()).Msg("Failed to start the server")
}
