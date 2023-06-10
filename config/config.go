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
		CacheConfig: CacheConfig{
			CacheDriver: getEnv("CACHE_DRIVER", "redis"),
			Redis: Redis{
				RedisHost:     getEnv("REDIS_HOST", "localhost"),
				RedisPort:     getEnvAsInt("REDIS_PORT", 6379),
				RedisDB:       getEnvAsInt("REDIS_DB", 0),
				RedisPassword: getEnv("REDIS_PASSWORD", ""),
			},
		},
	}
}

type Config struct {
	AppConfig
	DBconfig
	CacheConfig
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

type CacheConfig struct {
	CacheDriver string
	Redis       Redis
}
type Redis struct {
	RedisHost     string
	RedisPort     int
	RedisDB       int
	RedisPassword string
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
