package config

import (
	"os"
	"strconv"
)

func NewConfig() *Config {
	return &Config{
		AppConfig: AppConfig{
			AppHost: getEnv("APP_HOST", "localhost"),
			AppPort: getEnvAsInt("APP_PORT", 8000),
		},
		DBconfig: DBconfig{
			DBdriver:   getEnv("DB_DRIVER", ""),
			DBhost:     getEnv("DB_HOST", "localhost"),
			DBport:     getEnv("DB_PORT", "3306"),
			DBuser:     getEnv("DB_USER", "root"),
			DBpassword: getEnv("DB_PASSWORD", ""),
			DBname:     getEnv("DB_NAME", "go-simple-template"),
		},
	}
}

type Config struct {
	AppConfig
	DBconfig
}

type AppConfig struct {
	AppHost string
	AppPort int
}

type DBconfig struct {
	DBdriver   string
	DBhost     string
	DBport     string
	DBuser     string
	DBname     string
	DBpassword string
}

func getEnv(key string, defaultVal string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}

	if nextValue := os.Getenv(key); nextValue != "" {
		return nextValue
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
