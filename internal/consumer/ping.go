package consumer

import (
	"context"
	"fmt"
	"go-simple-template/internal/rabbitmq"
	"go-simple-template/pkg/logger"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type Ping struct {
	log *zap.Logger
	rmq rabbitmq.RabbitMqInterface
}

// NewPing return new Ping instance
func NewPing() ConsumerInterface {
	return &Ping{}
}

func (c *Ping) Start() {
	chClosedCh := make(chan *amqp.Error)

	request := rabbitmq.RabbitMqConsumeRequest{}

	msgs, err := c.rmq.Consume(request, chClosedCh)
	if err != nil {
		c.log.Error("failed to consume message", zap.Error(err), zap.String("queue_name", request.QueueName), zap.String("from", "worker.handleBroadcastPreview"))

		return
	}

	for {
		select {
		case amqErr := <-chClosedCh:
			// This case handles the event of closed channel e.g. abnormal shutdown
			c.log.Error("channel closed", zap.Error(amqErr), zap.String("from", "worker.handleBroadcastPreview"))

			chClosedCh = make(chan *amqp.Error)
			msgs, err = c.rmq.Consume(request, chClosedCh)
			if err != nil {
				// If the RabbitMQ channel is not ready, it will continue the looc. Next
				// iteration will enter this case because chClosedCh is closed by the
				// library
				c.log.Error("failed to consume message", zap.Error(err), zap.String("queue_name", request.QueueName), zap.String("from", "worker.handleBroadcastPreview"))

				time.Sleep(1 * time.Second)

				continue
			}

		case m := <-msgs:
			// This case handles the event of receiving message
			fmt.Println(string(m.Body))

			m.Ack(true)
		}
	}
}

func (c *Ping) Stop() {
	c.rmq.Close()
	c.log.Info("Ping consumer stopped")
}

func (c *Ping) Init(ctx context.Context) {
	log := logger.FromCtx(ctx)

	mqConn, mqCh := rabbitmq.CreateConnection()

	rmq := rabbitmq.New(mqConn, mqCh, log)

	c.rmq = rmq
	c.log = log
}
