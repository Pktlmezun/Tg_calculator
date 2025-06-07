package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	TelegramToken string
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")
	token := os.Getenv("TELEGRAM_TOKEN")
	return &Config{TelegramToken: token}
}
