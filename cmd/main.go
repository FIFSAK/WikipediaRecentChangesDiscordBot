package main

import (
	"WikipediaRecentChangesDiscordBot/bot"
	"WikipediaRecentChangesDiscordBot/config"
	"WikipediaRecentChangesDiscordBot/services/kafka"
	"WikipediaRecentChangesDiscordBot/services/redisClient"
	"WikipediaRecentChangesDiscordBot/services/wikipedia"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func main() {
	conf, err := config.New()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_ = kafka.NewKafka(conf)

	redis := redisClient.NewRedisClient(conf)

	var wg sync.WaitGroup

	wg.Add(1)
	go wikipedia.ListenToWikipediaChanges(&wg, redis)

	bot.Start(conf, redis)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		os.Exit(0)
	}()
	wg.Wait()
}
