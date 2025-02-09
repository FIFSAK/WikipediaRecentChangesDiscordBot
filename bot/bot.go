package bot

import (
	"WikipediaRecentChangesDiscordBot/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var (
	BotId string
)

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotId = u.ID

	goBot.AddHandler(recentHandler)
	goBot.AddHandler(setLangHandler)
	goBot.AddHandler(statsChangesHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}
