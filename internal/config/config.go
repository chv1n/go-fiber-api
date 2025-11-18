package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DBURL   string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		AppPort: os.Getenv("APP_PORT"),
		DBURL:   os.Getenv("DATABASE_URL"),
	}

	if cfg.AppPort == "" {
		log.Fatal("APP_PORT is required")
	}

	return cfg
}
