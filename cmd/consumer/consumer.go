package consumer

import (
	"context"
	"go-simple-template/internal/consumer"
)

func Start(ctx context.Context, consumer ...consumer.ConsumerInterface) {
	for _, c := range consumer {
		c.Init(ctx)
		c.Start()
	}
}
