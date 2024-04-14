package config

import "github.com/spf13/viper"

// AppPort retrieves the value of the APP_PORT environment variable.
func AppPort() int {
	return viper.GetInt("APP_PORT")
}

// AppHost retrieves the value of the APP_HOST environment variable.
func AppHost() string {
	return viper.GetString("APP_HOST")
}

// AppName retrieves the value of the APP_NAME environment variable.
func AppName() string {
	return viper.GetString("APP_NAME")
}

// AppEnv retrieves the value of the APP_ENV environment variable.
func AppEnv() string {
	return viper.GetString("APP_ENV")
}

// AppVersion retrieves the value of the APP_VERSION environment variable.
func AppVersion() string {
	return viper.GetString("APP_VERSION")
}
