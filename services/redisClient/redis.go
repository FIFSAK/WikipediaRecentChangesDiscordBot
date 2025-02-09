package redisClient

import (
	"WikipediaRecentChangesDiscordBot/config"
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
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

func IncrementChanges(date string, language string) {
	key := fmt.Sprintf("%s:%s", date, language)
	err := Rdb.Incr(Ctx, key).Err()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Printf("Incremented key: %s\n", key)

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
