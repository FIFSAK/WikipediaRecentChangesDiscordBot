package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Token         string
	BotPrefix     string
	RedisPassword string
	RedisAddr     string
)

func ReadConfig() error {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
	Token = os.Getenv("TOKEN")
	BotPrefix = os.Getenv("BOT_PREFIX")
	RedisPassword = os.Getenv("REDIS_PASSWORD")
	RedisAddr = os.Getenv("REDIS_HOST") + ":" + "6379"

	return nil
}
