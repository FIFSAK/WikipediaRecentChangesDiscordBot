package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Token     string
	BotPrefix string
)

func ReadConfig() error {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
	Token = os.Getenv("TOKEN")
	BotPrefix = os.Getenv("BOT_PREFIX")

	return nil
}
