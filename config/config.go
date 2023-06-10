package config

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func NewConfig(workDir string) *Config {
	err := os.Chdir(workDir)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed change to workspace directory")
		panic(err)
	}

	envPath := filepath.Join(workDir, ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Fatal().Err(err).Msg("Failed to load .env file")
		panic(err)
	}

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
