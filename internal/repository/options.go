package repository

import "gorm.io/gorm"

type Option func(*gorm.DB) *gorm.DB

// WithWhereClause is a function that takes a condition and returns a pointer to a gorm.DB.
// This allows us to chain multiple options together to build a query.
// This function can used for actions SELECT, UPDATE, DELETE.
func WithWhereClause(condition interface{}) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(condition)
	}
}

// WithSelectFields is a function that takes a list of fields and returns a pointer to a gorm.DB.
// This function is used to select only the fields that are needed.
// This function can be used for actions SELECT.
func WithSelectFields(fields ...string) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select(fields)
	}
}

// WithPagination is a function that takes a limit and offset and returns a pointer to a gorm.DB.
// This function is used to limit the number of records returned and to skip a number of records.
// This function can be used for actions SELECT.
func WithPagination(limit, offset int) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit).Offset(offset)
	}
}

// WithPreload is a function that takes a list of relations and returns a pointer to a gorm.DB.
// This function is used to preload relations.
// This function can be used for actions SELECT.
func WithPreload(relations ...string) Option {
	return func(db *gorm.DB) *gorm.DB {
		for _, relation := range relations {
			db = db.Preload(relation)
		}

		return db
	}
}

// WithOrder is a function that takes an order string and returns a pointer to a gorm.DB.
// This function is used to order the records.
// This function can be used for actions SELECT.
func WithOrder(order string) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}

// WithDeleted is a function that returns a pointer to a gorm.DB.
// This sets the global scopes to ignore soft deleted records.
// This function can be used for actions SELECT.
func WithDeleted() Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}
}

// WithCount is a function that takes a pointer to an int64 and returns a pointer to a gorm.DB.
// This function sets the count of records to the provided pointer.
// This function can be used for actions SELECT.
func WithCount(count *int64) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Count(count)
	}
}

// WithSoftDelete is a function that returns a pointer to a gorm.DB.
// This sets the global scopes to ignore soft deleted records.
// This is the opposite of WithDeleted.
// This function is automatically called when the model has gorm.DeletedAt field.
// This function can be used for actions SELECT.
func WithSoftDelete() Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("deleted_at IS NULL")
	}
}
