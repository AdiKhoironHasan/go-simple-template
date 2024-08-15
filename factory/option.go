package factory

import (
	"go-simple-template/internal/rabbitmq"
	"go-simple-template/pkg/cachex"
	"go-simple-template/pkg/storagex"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Option interface {
	apply(*Factory)
}

type optionFunc func(*Factory)

func (f optionFunc) apply(factory *Factory) {
	f(factory)
}

func WithDB(db *gorm.DB) Option {
	return optionFunc(func(factory *Factory) {
		factory.Db = db
	})
}

func WithStorage(storage *storagex.Storage) Option {
	return optionFunc(func(factory *Factory) {
		factory.Storage = storage
	})
}

func WithCache(cache *cachex.Cache) Option {
	return optionFunc(func(factory *Factory) {
		factory.Cache = cache
	})
}

func WithLogger(logger *zap.Logger) Option {
	return optionFunc(func(factory *Factory) {
		factory.Logger = logger
	})
}

func WithRabbitMQ(rmq rabbitmq.RabbitMqInterface) Option {
	return optionFunc(func(factory *Factory) {
		factory.RabbitMQ = rmq
	})
}
