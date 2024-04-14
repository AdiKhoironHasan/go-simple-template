package config

import "github.com/spf13/viper"

// RedisHost retrieves the value of the REDIS_HOST environment variable.
func RedisHost() string {
	return viper.GetString("REDIS_HOST")
}

// RedisPort retrieves the value of the REDIS_PORT environment variable.
func RedisPort() int {
	return viper.GetInt("REDIS_PORT")
}

// RedisDB retrieves the value of the REDIS_DB environment variable.
func RedisDB() int {
	return viper.GetInt("REDIS_DB")
}

// RedisPassword retrieves the value of the REDIS_PASSWORD environment variable.
func RedisPassword() string {
	return viper.GetString("REDIS_PASSWORD")
}
