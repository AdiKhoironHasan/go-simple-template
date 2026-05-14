package infrastructure

import "context"

func (f *Factory) BuildRestFactory(ctx context.Context) *Factory {
	f.setMongoDB(ctx)

	return f
}
