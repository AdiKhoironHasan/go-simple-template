package repository

import (
	"context"
)

type Storer interface {
	Store(ctx context.Context, input any) error
}

type Updater interface {
	Update(ctx context.Context, input any, opts ...Option) error
}

type Deleter interface {
	Delete(ctx context.Context, opts ...Option) error
}

type Counter interface {
	Count(ctx context.Context, opts ...Option) (int, error)
}
