package consumer

import "context"

type Consumer struct {
	Client ConsumerInterface
}

func New(consumer ConsumerInterface) *Consumer {
	return &Consumer{
		Client: consumer,
	}
}

type ConsumerInterface interface {
	Start()
	Stop()
	Init(ctx context.Context)
}
