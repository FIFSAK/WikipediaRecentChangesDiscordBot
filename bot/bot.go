package bot

import (
	"WikipediaRecentChangesDiscordBot/config"
	"WikipediaRecentChangesDiscordBot/services/redisClient"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var (
	BotId string
	Rc    *redisClient.RedisClient
)

func Start(c *config.Config, redis *redisClient.RedisClient) {
	goBot, err := discordgo.New("Bot " + c.Token)
	Rc = redis
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotId = u.ID

	handlers := NewHandlers(c)

	goBot.AddHandler(handlers.BotHandlers.recentHandler)
	goBot.AddHandler(handlers.BotHandlers.setLangHandler)
	goBot.AddHandler(handlers.BotHandlers.statsChangesHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}
