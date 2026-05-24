package rest

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go-simple-template/internal/adapter/inbound/rest/router"
	"go-simple-template/internal/adapter/inbound/rest/server"
	healthCache "go-simple-template/internal/adapter/outbound/cache/redis/health"
	tokenCacheAdapter "go-simple-template/internal/adapter/outbound/cache/redis/token"
	healthRepo "go-simple-template/internal/adapter/outbound/repository/mongo/health"
	userRepo "go-simple-template/internal/adapter/outbound/repository/mongo/user"
	"go-simple-template/internal/infrastructure"
	"go-simple-template/internal/pkg/config"
	authService "go-simple-template/internal/service/auth"
	healthService "go-simple-template/internal/service/health"
)

func Start(ctx context.Context) {
	//  Infrastructure
	mongo := infrastructure.NewMongoDB(ctx)
	err := mongo.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongo.Disconnect(ctx)

	redisClient := infrastructure.NewRedis(ctx)
	err = redisClient.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisClient.Disconnect(ctx)

	//  Outbound Adapters (Driven)
	healthRepo := healthRepo.New(mongo.Client)
	healthCache := healthCache.NewCache(redisClient.Client)
	tokenCache := tokenCacheAdapter.NewCache(redisClient.Client)
	userRepo := userRepo.New(mongo.Client, config.MongodbName())

	//  Application Services
	healthSvc := healthService.New(healthRepo, healthCache)
	authSvc := authService.New(userRepo, tokenCache)

	//  Inbound Adapters (Driving)
	deps := &router.Dependencies{
		HealthService: healthSvc,
		AuthService:   authSvc,
	}

	srv := server.New(ctx, deps)

	//  Graceful shutdown
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	//  Run Server
	if err := srv.Run(ctx); err != nil {
		log.Fatalf("failed to run rest server: %v", err)
	}
}
