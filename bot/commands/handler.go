package commands

import (
	"github.com/bwmarrin/discordgo"
)

func MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	switch m.Content {
	case "!ping":
		PingCommand(s, m)
	case "!help":
		HelpCommand(s, m)
	}
}
