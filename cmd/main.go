package main

import (
	"WikipediaRecentChangesDiscordBot/bot"
	"WikipediaRecentChangesDiscordBot/config"
	"WikipediaRecentChangesDiscordBot/listener"
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

	var wg sync.WaitGroup

	wg.Add(1)
	go listener.ListenToWikipediaChanges(&wg)

	bot.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		close(listener.LanguageFilterChan)
		os.Exit(0)
	}()
	wg.Wait()
}
