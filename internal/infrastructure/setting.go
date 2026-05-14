package infrastructure

import (
	"context"
	"log"
)

func (f *Factory) setMongoDB(ctx context.Context) *Factory {
	mongodb := NewMongoDB(ctx)
	err := mongodb.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// ping MongoDB to ensure connection is established
	err = mongodb.Client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	f.Mongodb = mongodb.Client
	return f
}

func (f *Factory) setRedis(ctx context.Context) *Factory {
	redis := NewRedis(ctx)
	redis.Connect(ctx)

	// ping Redis to ensure connection is established
	err := redis.Client.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	f.Redis = redis.Client
	return f
}
