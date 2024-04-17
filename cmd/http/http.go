package http

import (
	"context"
	"go-simple-template/factory"
	"go-simple-template/internal/database"
	"go-simple-template/internal/router"
	"go-simple-template/internal/server"
	"go-simple-template/pkg/cachex"
	"go-simple-template/pkg/cachex/redis"
	"go-simple-template/pkg/logger"
	"go-simple-template/pkg/storagex"
	"go-simple-template/pkg/storagex/minio"

	"go.uber.org/zap"
)

func Start(ctx context.Context) {
	log := logger.FromCtx(ctx)

	redis := redis.NewRedis()

	cache := cachex.NewCache(redis)

	db, err := database.NewConnection()
	if err != nil {
		log.Fatal("Failed to connect to database", zap.Error(err))
	}

	minio, err := minio.NewMinio()
	if err != nil {
		log.Fatal("Failed to connect to Minio", zap.Error(err))
	}

	storage := storagex.NewStorage(minio)

	factory := factory.New(
		factory.WithCache(cache),
		factory.WithDB(db),
		factory.WithStorage(storage),
		factory.WithLogger(log),
	)

	router := router.New(
		router.WithFactory(factory),
	)

	server := server.NewHttpServer(router)
	defer server.Done()

	if err := server.Run(ctx); err != nil {
		log.Fatal("Failed to start the server", zap.Error(err))
	}
}
