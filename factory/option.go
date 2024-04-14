package factory

import (
	"go-simple-template/pkg/cachex"
	"go-simple-template/pkg/storagex"

	"gorm.io/gorm"
)

type Option interface {
	apply(*Factory)
}

type optionFunc func(*Factory)

func (f optionFunc) apply(dep *Factory) {
	f(dep)
}

func WithDB(db *gorm.DB) Option {
	return optionFunc(func(dep *Factory) {
		dep.Db = db
	})
}

func WithStorage(storage *storagex.Storage) Option {
	return optionFunc(func(dep *Factory) {
		dep.Storage = storage
	})
}

func WithCache(cache *cachex.Cache) Option {
	return optionFunc(func(dep *Factory) {
		dep.Cache = cache
	})
}

// func WithQueue(queue *asynq.Client) Option {
// 	return optionFunc(func(dep *Factory) {
// 		dep.Queue = queue
// 	})
// }
