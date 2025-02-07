package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Token     string
	BotPrefix string
)

// ReadConfig reads the config.json file and unmarshals it into the Config struct
func ReadConfig() error {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
	Token = os.Getenv("TOKEN")
	BotPrefix = os.Getenv("BOT_PREFIX")
	fmt.Println("Token is:", Token+"\nBotPrefix is:", BotPrefix)

	return nil
}
