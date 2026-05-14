package health

import (
	"go-simple-template/internal/core/ports/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type health struct {
	client *mongo.Client
}

func New(client *mongo.Client) repository.HealthRepository {
	return &health{
		client: client,
	}
}
