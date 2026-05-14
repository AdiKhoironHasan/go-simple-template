package rest

import (
	"context"
	"log"

	"go-simple-template/internal/adapter/inbound/rest/router"
	"go-simple-template/internal/adapter/inbound/rest/server"
	healthMongoRepo "go-simple-template/internal/adapter/outbound/mongo/health"
	healthCacheRepo "go-simple-template/internal/adapter/outbound/redis/health"
	healthService "go-simple-template/internal/app/health"
	"go-simple-template/internal/infrastructure"
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
	healthRepo := healthMongoRepo.New(mongo.Client)
	healthCacheRepo := healthCacheRepo.NewCache(redis.Client)

	//  Application Services
	healthSvc := healthService.New(healthRepo, healthCacheRepo)

	//  Inbound Adapters (Driving)
	deps := &router.Dependencies{
		HealthService: healthSvc,
	}

	srv := server.New(ctx, deps)

	//  Run Server
	if err := srv.Run(ctx); err != nil {
		log.Fatalf("failed to run rest server: %v", err)
	}
}
