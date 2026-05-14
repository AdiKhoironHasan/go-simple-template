package config

import "github.com/spf13/viper"

func MongodbProtocol() string {
	return viper.GetString("MONGODB_PROTOCOL")
}

func MongodbAddress() string {
	return viper.GetString("MONGODB_ADDRESS")
}

func MongodbUsername() string {
	return viper.GetString("MONGODB_USERNAME")
}

func MongodbPassword() string {
	return viper.GetString("MONGODB_PASSWORD")
}

func MongodbMaxConnOpen() int {
	return viper.GetInt("MONGODB_MAX_CONN_OPEN")
}

func MongodbMaxConnIdle() int {
	return viper.GetInt("MONGODB_MAX_CONN_IDLE")
}

func MongodbMaxConnLifetime() string {
	return viper.GetString("MONGODB_MAX_CONN_LIFETIME")
}

func MongodbOption() string {
	return viper.GetString("MONGODB_OPTION")
}

func MongodbName() string {
	return viper.GetString("MONGODB_NAME")
}
