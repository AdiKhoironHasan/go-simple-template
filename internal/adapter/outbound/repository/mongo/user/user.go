package user

import (
	"github.com/adikhoironhasan/go-simple-template/internal/adapter/outbound/repository/mongo/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type user struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func New(client *mongo.Client, dbName string) *user {
	return &user{
		client:     client,
		collection: client.Database(dbName).Collection(model.User{}.CollectionName()),
	}
}
