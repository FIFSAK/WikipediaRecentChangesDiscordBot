package bot

import (
	"WikipediaRecentChangesDiscordBot/config"
	"WikipediaRecentChangesDiscordBot/listener"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var (
	BotId string
	//allowedLanguages = map[string]string{"en": "enwiki", "de": "dewiki", "fr": "frwiki", "es": "eswiki", "it": "itwiki", "ru": "ruwiki"}
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

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}

func recentHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == config.BotPrefix+"recent" {
		fmt.Println("Recent command received")
		recent := listener.GetRecentChanges()
		fmt.Println("Recent changes: ", recent)
		if len(recent) == 0 {
			_, err := s.ChannelMessageSend(m.ChannelID, "No recent changes available.")
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		var message string
		cnt := 1
		for _, r := range recent {
			if len(message)+len(fmt.Sprintf("%d.\n"+r.String()+"\n", cnt)) <= 2000 {
				message += fmt.Sprintf("%d.\n"+r.String()+"\n", cnt)
				cnt++

			} else {
				break
			}
		}
		_, err := s.ChannelMessageSend(m.ChannelID, message)
		if err != nil {
			fmt.Println(err.Error())

		}
	}

}

func setLangHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, config.BotPrefix+"setLang") {
		language := strings.TrimSpace(strings.TrimPrefix(m.Content, config.BotPrefix+"setLang "))
		if language == "" || language == config.BotPrefix+"setLang" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Please specify a language code, e.g., !setLang en")
			return
		}

		allowedLanguages := map[string]string{
			"en": "enwiki",
			"de": "dewiki",
			"fr": "frwiki",
			"es": "eswiki",
			"it": "itwiki",
			"ru": "ruwiki",
		}

		if _, ok := allowedLanguages[language]; !ok {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Language not supported. Supported languages: en, de, fr, es, it, ru")
			return
		}

		select {
		case listener.LanguageFilterChan <- allowedLanguages[language]:
			_, _ = s.ChannelMessageSend(m.ChannelID, "Language updated to: "+language)

		default:
			_, _ = s.ChannelMessageSend(m.ChannelID, "Error: Unable to update language. Please try again.")
		}
	}
}
