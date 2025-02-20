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
	conf, _ := config.New()
	kafkaConn := kafka.NewKafka(conf)
	defer kafkaConn.Close()

	redisConn := redisClient.NewRedisClient(conf, kafkaConn)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		wikipedia.ListenToWikipediaChanges(kafkaConn)
	}()

	go redisConn.ConsumeChangesFromKafka(&wg)

	bot.Start(conf, redisConn)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		fmt.Println("Got interrupt signal, exiting...")
		os.Exit(0)
	}()
	wg.Wait()
}
