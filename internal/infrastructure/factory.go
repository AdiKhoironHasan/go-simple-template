package infrastructure

import (
	red "github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Factory struct {
	Mongodb *mongo.Client
	Redis   *red.Client
}

func NewFactory() *Factory {
	return &Factory{}
}
