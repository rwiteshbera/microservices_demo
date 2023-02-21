package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	DB_DRIVER            string
	DB_CONNECTION_STRING string
	SERVER_HOST          string
	SERVER_PORT          string
}

func LoadEnvConfig() (*EnvConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	config := &EnvConfig{
		DB_DRIVER:            os.Getenv("DB_DRIVER"),
		DB_CONNECTION_STRING: os.Getenv("DB_CONNECTION_STRING"),
		SERVER_HOST:          os.Getenv("SERVER_HOST"),
		SERVER_PORT:          os.Getenv("SERVER_PORT"),
	}

	return config, nil
}
