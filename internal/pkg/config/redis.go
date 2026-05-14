package config

import "github.com/spf13/viper"

func RedisHost() string {
	return viper.GetString("REDIS_HOST")
}

func RedisPort() int {
	return viper.GetInt("REDIS_PORT")
}

func RedisUsername() string {
	return viper.GetString("REDIS_USERNAME")
}

func RedisPassword() string {
	return viper.GetString("REDIS_PASSWORD")
}

func RedisDB() int {
	return viper.GetInt("REDIS_DB")
}

func RedisTLS() bool {
	return viper.GetBool("REDIS_TLS")
}
