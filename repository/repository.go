package repository

import (
	"go-simple-template/pkg/cachex"
	"go-simple-template/pkg/logger"

	"gorm.io/gorm"
)

type repository struct {
	db    *gorm.DB
	cache *cachex.Cache
}

var (
	logRepo = logger.NewLogger().Logger.With().Str("pkg", "repository").Logger()
)

func NewRepository() *repository {
	return &repository{}
}

type RepositoryInterface interface {
	Ping() error
}

func (r *repository) WithDB(db *gorm.DB) *repository {
	r.db = db
	return r
}

func (r *repository) WithCache(cache *cachex.Cache) *repository {
	r.cache = cache
	return r
}
