package redisClient

import (
	"WikipediaRecentChangesDiscordBot/config"
	kfk "WikipediaRecentChangesDiscordBot/services/kafka"
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var Kc *kfk.Kafka

type RedisClient struct {
	Client *redis.Client
	Config *config.Config
	Ctx    context.Context
	Kafka  *kfk.Kafka
}

func NewRedisClient(config *config.Config) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       0,
	})
	Kafka := kfk.NewKafka(config)
	return &RedisClient{
		Client: rdb,
		Config: config,
		Ctx:    context.Background(),
		Kafka:  Kafka,
	}

}

func (r *RedisClient) IncrementChanges(date string, language string) error {
	for {
		msg, err := r.Kafka.Reader.ReadMessage(context.Background())

		if err != nil {
			log.Fatalf("failed to read message: %v", err)
		}
		fmt.Println(string(msg.Value))
		key := string(msg.Value)

		exists, err := r.Client.Exists(r.Ctx, key).Result()
		if err != nil {
			return fmt.Errorf("failed to check existence: %w", err)
		}

		if exists == 0 {
			err = r.Client.Set(r.Ctx, key, 1, time.Hour*24*7).Err()
			if err != nil {
				return fmt.Errorf("failed to set key: %w", err)
			}
		} else {
			_, err = r.Client.Incr(r.Ctx, key).Result()
			if err != nil {
				return fmt.Errorf("failed to incr key: %w", err)
			}
		}

	}
	return nil
}

func (r *RedisClient) GetChanges(date string, language string) int {
	key := fmt.Sprintf("%s:%s", date, language)
	value, err := r.Client.Get(r.Ctx, key).Int()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			fmt.Printf("No data for key: %s\n", key)
			return 0
		}
		fmt.Printf("Failed to get key %s: %v\n", key, err)
		return 0
	}
	return value
}
