package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Token         string
	BotPrefix     string
	RedisPassword string
	RedisAddr     string
	KafkaBroker   string
	KafkaTopic    string
	KafkaGroup    string
}

func New() (*Config, error) {
	c := Config{}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
		return nil, err
	}
	c.Token = os.Getenv("TOKEN")
	c.BotPrefix = os.Getenv("BOT_PREFIX")
	c.RedisPassword = os.Getenv("REDIS_PASSWORD")
	c.RedisAddr = os.Getenv("REDIS_HOST") + ":" + "6379"
	c.KafkaBroker = os.Getenv("KAFKA_BROKER")
	c.KafkaTopic = os.Getenv("KAFKA_TOPIC")
	c.KafkaGroup = os.Getenv("KAFKA_GROUP")
	return &c, nil
}
