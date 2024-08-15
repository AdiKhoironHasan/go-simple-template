package rabbitmq

import (
	"fmt"
	"go-simple-template/config"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

type (
	mqConfig struct {
		User     string
		Password string
		VHost    string
		Host     string
		Port     int
	}

	rabbitMqConn struct {
		mqConfig
	}
)

var (
	rmqConn *amqp.Connection
	rmqChan *amqp.Channel
	rmqConf mqConfig
	once    sync.Once
)

func CreateConnection() (*amqp.Connection, *amqp.Channel) {
	once.Do(func() {
		rmqConf = mqConfig{
			Host:     config.RabbitMQHost(),
			Port:     config.RabbitMQPort(),
			User:     config.RabbitMQUser(),
			Password: config.RabbitMQPassword(),
			VHost:    config.RabbitMQVHost(),
		}
	})

	rabbitMqConn := rabbitMqConn{mqConfig: rmqConf}
	if (rmqConn == nil && rmqChan == nil) || rmqChan.IsClosed() || rmqConn.IsClosed() {
		rmqConn, rmqChan = rabbitMqConn.connect()
	}

	return rmqConn, rmqChan
}

func (conf rabbitMqConn) connect() (*amqp.Connection, *amqp.Channel) {
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.VHost)
	conn, err := amqp.Dial(connStr)
	if err != nil {
		panic(err)
	}

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return conn, channel
}
