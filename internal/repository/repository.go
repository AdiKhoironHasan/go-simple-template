package repository

import (
	"context"
	"go-simple-template/pkg/cachex"

	"gorm.io/gorm"
)

type repository struct {
	db    *gorm.DB
	cache *cachex.Cache
}

func NewPing(db *gorm.DB, cache *cachex.Cache) PingRepository {
	return &repository{
		db:    db,
		cache: cache,
	}
}

type PingRepository interface {
	Ping(ctx context.Context) error
}
