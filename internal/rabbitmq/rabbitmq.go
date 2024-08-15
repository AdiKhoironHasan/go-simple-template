package rabbitmq

import (
	"context"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type RabbitMqInterface interface {
	Consume(request RabbitMqConsumeRequest, chClosedCh chan *amqp.Error) (<-chan amqp.Delivery, error)
	Publish(ctx context.Context, request RabbitMqPublishRequest) error
	DeclareExchange(exchange RabbitMQExchange)
	BindingQueue(exchangeName string, queueName string)
	DeclareQueue(queueName string) amqp.Queue
	InspectQueue(queueName string) (*amqp.Queue, error)
	DeleteQueue(queueName string) error
	Purge(queueName string) error
	Close()
	reconnect()
	IsClosed() bool
}

type rabbitMq struct {
	mqConn   *amqp.Connection
	mqCh     *amqp.Channel
	shutdown bool
	log      *zap.Logger
}

func New(conn *amqp.Connection, ch *amqp.Channel, log *zap.Logger) RabbitMqInterface {
	return &rabbitMq{
		mqConn: conn,
		mqCh:   ch,
		log:    log,
	}
}

func (r *rabbitMq) reconnect() {
	if r.shutdown {
		r.log.Error("rabbitmq is shutdown by calling Close()")

		return
	}

	if r.mqConn != nil {
		newCh, err := r.mqConn.Channel()
		if err == nil {
			r.log.Warn("reconnect rabbitmq by channel")
			r.mqCh = newCh
			return
		}
	}

	r.log.Warn("reconnect rabbitmq by connection")
	r.mqConn, r.mqCh = CreateConnection()
}

func (r *rabbitMq) Publish(ctx context.Context, request RabbitMqPublishRequest) error {
	if r.mqCh.IsClosed() || r.mqConn.IsClosed() {
		r.reconnect()
	}

	cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	body, err := json.Marshal(request.Messages)
	if err != nil {
		r.log.Error("failed to marshal message", zap.Error(err))

		return err
	}

	publish := amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
		Headers:     request.Headers,
	}

	err = r.mqCh.PublishWithContext(cctx,
		request.Exchange,  // exchange
		request.QueueName, // routing key
		true,              // mandatory
		false,             // immediate
		publish,
	)

	if err != nil {
		r.log.Error("failed to publish message", zap.Error(err))
		return err
	}

	return nil
}

func (r *rabbitMq) DeleteQueue(queueName string) error {
	if r.mqCh.IsClosed() || r.mqConn.IsClosed() {
		r.reconnect()
	}

	_, err := r.mqCh.QueueDelete(queueName, false, false, false)
	if err != nil {
		r.log.Error("failed to delete queue", zap.Error(err))

		return err
	}

	return nil
}

func (r *rabbitMq) Consume(request RabbitMqConsumeRequest, chClosedCh chan *amqp.Error) (<-chan amqp.Delivery, error) {
	if r.mqCh.IsClosed() || r.mqConn.IsClosed() {
		r.reconnect()
	}

	if err := r.mqCh.Qos(
		1, // prefetch count, 1 means process one message at a time
		0, // prefetch size, 0 means no limit
		false,
	); err != nil {
		return nil, err
	}

	msgs, err := r.mqCh.Consume(
		request.QueueName,    // queue
		request.ConsumerName, // consumer
		false,                // auto-ack
		false,                // exclusive
		false,                // no-local
		false,                // no-wait
		nil,                  // args
	)

	if err != nil {
		r.log.Error("failed to register consumer", zap.Error(err))

		return nil, err
	}

	// notify when channel closed
	r.mqCh.NotifyClose(chClosedCh)

	return msgs, nil
}

func (r *rabbitMq) DeclareExchange(exchange RabbitMQExchange) {
	if r.mqCh.IsClosed() || r.mqConn.IsClosed() {
		r.reconnect()
	}

	err := r.mqCh.ExchangeDeclare(
		exchange.Name,
		exchange.Kind,
		true,
		false,
		false,
		false,
		exchange.Args,
	)

	if err != nil {
		r.log.Error("failed to declare exchange", zap.Error(err))
	}
}

func (r *rabbitMq) BindingQueue(exchangeName string, queueName string) {
	if r.mqCh.IsClosed() || r.mqConn.IsClosed() {
		r.reconnect()
	}

	r.DeclareQueue(queueName)

	err := r.mqCh.QueueBind(queueName, "", exchangeName, false, nil)
	if err != nil {
		r.log.Error("failed to bind queue to exchange", zap.Error(err), zap.String("queue", queueName), zap.String("exchange", exchangeName))
	}
}

func (r *rabbitMq) DeclareQueue(queueName string) amqp.Queue {
	if r.mqCh.IsClosed() || r.mqConn.IsClosed() {
		r.reconnect()
	}

	q, err := r.mqCh.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		r.log.Error("failed to declare queue", zap.Error(err), zap.String("queue", queueName))
	}

	return q
}

func (r *rabbitMq) InspectQueue(queueName string) (*amqp.Queue, error) {
	if r.mqCh.IsClosed() || r.mqConn.IsClosed() {
		r.reconnect()
	}

	queue, err := r.mqCh.QueueInspect(queueName)
	if err != nil {
		r.log.Error("failed to inspect queue", zap.Error(err), zap.String("queue", queueName))

		return nil, err
	}

	return &queue, nil
}

func (r *rabbitMq) Purge(queueName string) error {
	if r.mqCh.IsClosed() || r.mqConn.IsClosed() {
		r.reconnect()
	}

	_, err := r.mqCh.QueuePurge(queueName, false)
	if err != nil {
		r.log.Error("failed to purge queue", zap.Error(err), zap.String("queue", queueName))

		return err
	}

	return nil
}

func (r *rabbitMq) Close() {
	if !r.mqCh.IsClosed() {
		r.log.Warn("Closing rabbitMq channel...")
		r.mqCh.Close()
	}

	if !r.mqConn.IsClosed() {
		r.log.Warn("Closing rabbitMq connection...")
		r.mqConn.Close()
	}

	r.shutdown = true

	r.log.Warn("RabbitMq connection closed")
}

func (r *rabbitMq) IsClosed() bool {
	return r.shutdown
}
