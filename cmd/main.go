package main

import (
	"WikipediaRecentChangesDiscordBot/bot"
	"WikipediaRecentChangesDiscordBot/config"
	"WikipediaRecentChangesDiscordBot/services/redisClient"
	"WikipediaRecentChangesDiscordBot/services/wikipedia"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	redisClient.InitializeRedis()

	var wg sync.WaitGroup

	wg.Add(1)
	go wikipedia.ListenToWikipediaChanges(&wg)

	bot.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		os.Exit(0)
	}()
	wg.Wait()
}
