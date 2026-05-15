package rest

import (
	"context"
	"log"

	"go-simple-template/internal/adapter/inbound/rest/router"
	"go-simple-template/internal/adapter/inbound/rest/server"
	healthRepo "go-simple-template/internal/adapter/outbound/repository/mongo/health"
	healthCache "go-simple-template/internal/adapter/outbound/cache/redis/health"
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
	healthRepo := healthRepo.New(mongo.Client)
	healthCache := healthCache.NewCache(redis.Client)

	//  Application Services
	healthSvc := healthService.New(healthRepo, healthCache)

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
