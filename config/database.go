package config

import "github.com/spf13/viper"

// DBDriver retrieves the value of the DB_DRIVER environment variable.
func DBDriver() string {
	return viper.GetString("DB_DRIVER")
}

// DBHost retrieves the value of the DB_HOST environment variable.
func DBHost() string {
	return viper.GetString("DB_HOST")
}

// DBPort retrieves the value of the DB_PORT environment variable.
func DBPort() int {
	return viper.GetInt("DB_PORT")
}

// DBUser retrieves the value of the DB_USER environment variable.
func DBUser() string {
	return viper.GetString("DB_USER")
}

// DBPassword retrieves the value of the DB_PASSWORD environment variable.
func DBPassword() string {
	return viper.GetString("DB_PASSWORD")
}

// DBName retrieves the value of the DB_NAME environment variable.
func DBName() string {
	return viper.GetString("DB_NAME")
}
