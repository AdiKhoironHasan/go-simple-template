package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type (
	RabbitMQExchange struct {
		Name string
		Kind string
		Args amqp.Table
	}

	RabbitMqPublishRequest struct {
		Exchange  string
		QueueName string
		Headers   amqp.Table
		Messages  interface{}
	}

	RabbitMqConsumeRequest struct {
		QueueName    string
		ConsumerName string
	}

	Preparation struct {
		IsBindingExchange bool
		Exchange          RabbitMQExchange
		QueueName         string
	}
)
