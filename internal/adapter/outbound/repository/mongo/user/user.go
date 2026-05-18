package user

import (
	"go-simple-template/internal/adapter/outbound/repository/mongo/model"
	"go-simple-template/internal/core/port/outbound/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type user struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func New(client *mongo.Client, dbName string) repository.UserRepository {
	return &user{
		client:     client,
		collection: client.Database(dbName).Collection(model.User{}.CollectionName()),
	}
}
