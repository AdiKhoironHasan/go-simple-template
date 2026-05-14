package health

import (
	"go-simple-template/internal/core/port/outbound"

	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	client *mongo.Client
}

func New(client *mongo.Client) outbound.HealthRepository {
	return &repository{
		client: client,
	}
}
