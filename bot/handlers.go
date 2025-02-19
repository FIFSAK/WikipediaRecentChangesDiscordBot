package bot

import "WikipediaRecentChangesDiscordBot/config"

type Handlers struct {
	BotHandlers *BotHandlers
}

func NewHandlers(config *config.Config) *Handlers {
	return &Handlers{
		BotHandlers: &BotHandlers{config},
	}
}
