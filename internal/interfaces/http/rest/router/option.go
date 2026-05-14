package router

import "go-simple-template/internal/infrastructure"

type Option func(*router)

func WithFactory(f *infrastructure.Factory) Option {
	return func(r *router) {
		r.factory = f
	}
}
