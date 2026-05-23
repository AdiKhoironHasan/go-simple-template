package rest

import (
	"context"
	"log"

	"go-simple-template/internal/adapter/inbound/rest/router"
	"go-simple-template/internal/adapter/inbound/rest/server"
	healthCache "go-simple-template/internal/adapter/outbound/cache/redis/health"
	tokenCache "go-simple-template/internal/adapter/outbound/cache/redis/token"
	healthRepo "go-simple-template/internal/adapter/outbound/repository/mongo/health"
	healthService "go-simple-template/internal/app/health"
	"go-simple-template/internal/infrastructure"

	userRepo "go-simple-template/internal/adapter/outbound/repository/mongo/user"
	authService "go-simple-template/internal/app/auth"
	"go-simple-template/internal/pkg/config"
)

func Start(ctx context.Context) {
	//  Infrastructure
	mongo := infrastructure.NewMongoDB(ctx)
	err := mongo.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	redis := infrastructure.NewRedis(ctx)
	redis.Connect(ctx)

	//  Outbound Adapters (Driven)
	healthRepo := healthRepo.New(mongo.Client)
	healthCache := healthCache.NewCache(redis.Client)
	tokenCache := tokenCache.NewCache(redis.Client)
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

	//  Run Server
	if err := srv.Run(ctx); err != nil {
		log.Fatalf("failed to run rest server: %v", err)
	}
}
