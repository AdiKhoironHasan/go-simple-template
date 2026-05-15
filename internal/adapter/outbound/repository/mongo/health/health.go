package health

import (
	"go-simple-template/internal/core/port/outbound/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	client *mongo.Client
}

func New(client *mongo.Client) repository.Health {
	return &repo{
		client: client,
	}
}
