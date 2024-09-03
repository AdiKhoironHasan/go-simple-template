package repository

import (
	"context"
)

type Storer interface {
	Store(ctx context.Context, input any) error
}

type Updater interface {
	Update(ctx context.Context, input any, opts ...UpdateOption) error
}

type Deleter interface {
	Delete(ctx context.Context, id uint, opts ...DeleteOption) error
}

type Counter interface {
	Count(ctx context.Context, opts ...FindOption) (int, error)
}
