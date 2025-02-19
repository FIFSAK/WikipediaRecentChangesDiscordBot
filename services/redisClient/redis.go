package redisClient

import (
	"WikipediaRecentChangesDiscordBot/config"
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var Rdb *redis.Client
var Ctx = context.Background()

func InitializeRedis() {

	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       0,
	})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		fmt.Printf(err.Error())
		panic(err)
	}

	fmt.Println("Successfully connected to Redis")
}

func IncrementChanges(date string, language string) error {
	key := fmt.Sprintf("%s:%s", date, language)

	exists, err := Rdb.Exists(Ctx, key).Result()
	if err != nil {
		return fmt.Errorf("failed to check existence: %w", err)
	}

	if exists == 0 {
		err = Rdb.Set(Ctx, key, 1, time.Hour*24*7).Err()
		if err != nil {
			return fmt.Errorf("failed to set key: %w", err)
		}
	} else {
		_, err = Rdb.Incr(Ctx, key).Result()
		if err != nil {
			return fmt.Errorf("failed to incr key: %w", err)
		}
	}

	return nil
}

func GetChanges(date string, language string) int {
	key := fmt.Sprintf("%s:%s", date, language)
	value, err := Rdb.Get(Ctx, key).Int()
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
