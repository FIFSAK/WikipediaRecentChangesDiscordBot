package redisClient

import (
	"WikipediaRecentChangesDiscordBot/config"
	"WikipediaRecentChangesDiscordBot/services/kafka"
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
	"time"
)

type RedisClient struct {
	Client *redis.Client
	Config *config.Config
	Ctx    context.Context
	Kafka  *kafka.Kafka
}

func NewRedisClient(config *config.Config, kafkaConn *kafka.Kafka) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       0,
	})
	return &RedisClient{
		Client: rdb,
		Config: config,
		Ctx:    context.Background(),
		Kafka:  kafkaConn,
	}

}

func (r *RedisClient) ConsumeChangesFromKafka(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		msg, err := r.Kafka.Reader.ReadMessage(r.Ctx)
		if err != nil {
			log.Printf("failed to read message: %v", err)
			return
		}

		key := string(msg.Value)
		exists, err := r.Client.Exists(r.Ctx, key).Result()
		if err != nil {
			log.Printf("failed to check existence: %v", err)
			continue
		}
		if exists == 0 {
			err = r.Client.Set(r.Ctx, key, 1, 24*time.Hour*7).Err()
			if err != nil {
				log.Printf("failed to set key: %v", err)
				continue
			}
		} else {
			_, err = r.Client.Incr(r.Ctx, key).Result()
			if err != nil {
				log.Printf("failed to incr key: %v", err)
				continue
			}
		}
	}
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
