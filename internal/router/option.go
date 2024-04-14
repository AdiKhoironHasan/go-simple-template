package router

import "go-simple-template/factory"

// Option return Router with RouterOption to fill up the dependencies from factory
type Option func(*Router)

// WithFactory is an option
func WithFactory(factory *factory.Factory) Option {
	return func(r *Router) {
		r.Factory = factory
	}
}
