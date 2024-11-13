package repository

import (
	"context"
	"go-simple-template/internal/repository/option"
)

type Storer interface {
	Store(ctx context.Context, input any, opts ...option.Option) error
}

type Updater interface {
	Update(ctx context.Context, input any, opts ...option.Option) error
}

type Deleter interface {
	Delete(ctx context.Context, opts ...option.Option) error
}

type Counter interface {
	Count(ctx context.Context, opts ...option.Option) (int, error)
}
