package factory

import (
	"go-simple-template/pkg/cachex"
	"go-simple-template/pkg/storagex"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Factory struct {
	Db      *gorm.DB
	Storage *storagex.Storage
	Cache   *cachex.Cache
	Logger  *zap.Logger
}

func New(opts ...Option) *Factory {
	dep := &Factory{}

	for _, opt := range opts {
		opt.apply(dep)
	}

	return dep
}
