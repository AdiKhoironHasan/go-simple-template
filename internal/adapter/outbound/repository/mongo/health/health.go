package health

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	client *mongo.Client
}

func New(client *mongo.Client) *repo {
	return &repo{
		client: client,
	}
}
