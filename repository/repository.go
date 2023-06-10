package repository

import (
	"go-simple-template/cache"
	"go-simple-template/pkg/logger"

	"gorm.io/gorm"
)

type repository struct {
	db    *gorm.DB
	cache *cache.Cache
}

var (
	logRepo = logger.NewLogger().Logger.With().Str("pkg", "repository").Logger()
)

func NewRepository(db *gorm.DB, cache *cache.Cache) *repository {
	return &repository{
		db:    db,
		cache: cache,
	}
}

type RepositoryInterface interface {
	Ping() error
}
