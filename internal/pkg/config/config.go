package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadEnv(path string) {
	// Load .env file
	if err := godotenv.Load(path); err != nil {
		log.Println("No .env file found, using os environment variables")
	}

	// Set Viper to read from environment variables
	viper.AutomaticEnv()
}
