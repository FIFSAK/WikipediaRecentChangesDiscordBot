package bot

import (
	"WikipediaRecentChangesDiscordBot/config"
	"WikipediaRecentChangesDiscordBot/services/wikipedia"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

var language string

type BotHandlers struct {
	Config *config.Config
}

func (h BotHandlers) recentHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if BotId == m.Author.ID {
		return
	}
	if m.Content == h.Config.BotPrefix+"recent" {
		recent := wikipedia.GetRecentChanges()
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

func (h BotHandlers) setLangHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if BotId == m.Author.ID {
		return
	}
	if strings.HasPrefix(m.Content, h.Config.BotPrefix+"setLang") {
		language = strings.TrimSpace(strings.TrimPrefix(m.Content, h.Config.BotPrefix+"setLang "))
		if language == "" || language == h.Config.BotPrefix+"setLang" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Please specify a language code, e.g., !setLang en")
			return
		}

		if _, ok := wikipedia.AllowedLanguages[language]; !ok {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Language not supported. Supported languages: en, de, fr, es, it, ru, any to all languages")
			return
		}

		select {
		case wikipedia.LanguageFilterChan <- wikipedia.AllowedLanguages[language]:
			_, _ = s.ChannelMessageSend(m.ChannelID, "Language updated to: "+language)

		default:
			_, _ = s.ChannelMessageSend(m.ChannelID, "Error: Unable to update language. Please try again.")
		}
	}
}

func (h BotHandlers) statsChangesHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if BotId == m.Author.ID {
		return
	}
	if strings.HasPrefix(m.Content, h.Config.BotPrefix+"stats") {
		args := strings.Split(m.Content, " ")
		if len(args) != 2 {
			_, err := s.ChannelMessageSend(m.ChannelID, "Usage: !stats [yyyy-mm-dd]")
			if err != nil {
				fmt.Println(err.Error())
			}
			return
		}

		date := args[1]

		if _, err := time.Parse("2006-01-02", date); err != nil {
			_, err = s.ChannelMessageSend(m.ChannelID, "Invalid date format. Use YYYY-MM-DD.")
			if err != nil {
				fmt.Println(err.Error())
			}
			return
		}
		if language == "" || language == "any" {
			_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Set a specific language by using !setLang"))
			if err != nil {
				fmt.Println(err.Error())
			}
			return
		}
		changes := Rc.GetChanges(date, wikipedia.AllowedLanguages[language])
		if changes == 0 {
			_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("No changes found for %s on %s.", language, date))
			if err != nil {
				fmt.Println(err.Error())
			}
			return
		}

		message := fmt.Sprintf("On %s, there were %d changes in %s.", date, changes, language)

		_, err := s.ChannelMessageSend(m.ChannelID, message)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
