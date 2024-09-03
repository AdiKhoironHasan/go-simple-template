package repository

import "gorm.io/gorm"

type FindOption func(*gorm.DB) *gorm.DB

func WithWhereClause(condition interface{}) FindOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(condition)
	}
}

func WithSelectFields(fields ...string) FindOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select(fields)
	}
}

func WithPagination(limit, offset int) FindOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit).Offset(offset)
	}
}

func WithPreload(relations ...string) FindOption {
	return func(db *gorm.DB) *gorm.DB {
		for _, relation := range relations {
			db = db.Preload(relation)
		}

		return db
	}
}

func WithOrder(order string) FindOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}

func WithDeleted() FindOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}
}

func WithCount(count *int64) FindOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Count(count)
	}
}

type UpdateOption func(*gorm.DB) *gorm.DB

func WithUpdateFields(fields map[string]interface{}) UpdateOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Updates(fields)
	}
}

type DeleteOption func(*gorm.DB) *gorm.DB

func WithSoftDelete() DeleteOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("deleted_at IS NULL")
	}
}
