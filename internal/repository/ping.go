package repository

import (
	"context"
	"go-simple-template/pkg/tracer"

	"github.com/stretchr/testify/assert"
)

func (r *repository) Ping(ctx context.Context) error {
	ctx, span := tracer.SpanStart(ctx, "Repository.Ping")
	defer span.Finish()

	_, err := r.cache.Client.Ping(ctx)
	err = assert.AnError
	if err != nil {
		span.AddError(err)

		return err
	}

	if err := r.db.Exec("SELECT 1").Error; err != nil {
		span.AddError(err)

		return err
	}

	return nil
}
