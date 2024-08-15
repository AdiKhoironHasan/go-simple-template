package config

import "github.com/spf13/viper"

func RabbitMQHost() string {
	return viper.GetString("RABBITMQ_HOST")
}

func RabbitMQPort() int {
	return viper.GetInt("RABBITMQ_PORT")
}

func RabbitMQUser() string {
	return viper.GetString("RABBITMQ_USER")
}

func RabbitMQPassword() string {
	return viper.GetString("RABBITMQ_PASSWORD")
}

func RabbitMQVHost() string {
	return viper.GetString("RABBITMQ_VHOST")
}
