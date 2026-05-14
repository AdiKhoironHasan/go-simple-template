package infrastructure

import (
	"context"
	"fmt"
	"log/slog"

	"go-simple-template/internal/pkg/config"
	"go-simple-template/internal/pkg/consts"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongodb struct {
	connectionURL string
	Client        *mongo.Client
}

func NewMongoDB(ctx context.Context) *mongodb {
	return &mongodb{
		connectionURL: fmt.Sprintf("%s://%s:%s@%s/%s%s", config.MongodbProtocol(), config.MongodbUsername(), config.MongodbPassword(), config.MongodbAddress(), config.MongodbName(), config.MongodbOption()),
	}
}

func (a *mongodb) Connect(ctx context.Context) error {
	clientOptions := options.Client().ApplyURI(a.connectionURL)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to connect to MongoDB", slog.String("connection", a.connectionURL), slog.String(consts.Error, err.Error()))
		return err
	}

	a.Client = client

	err = client.Ping(ctx, nil)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to ping MongoDB", slog.String("connection", a.connectionURL), slog.String(consts.Error, err.Error()))
		return err
	}

	slog.InfoContext(ctx, "MongoDB connected", slog.String("connection", a.connectionURL))

	return nil
}

func (a *mongodb) Disconnect(ctx context.Context) {
	err := a.Client.Disconnect(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to disconnect to MongoDB", slog.String("connection", a.connectionURL), slog.String(consts.Error, err.Error()))
	}
}
