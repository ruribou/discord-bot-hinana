package commands

import (
	"github.com/bwmarrin/discordgo"
)

func HelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "`!ping`: Pong! を返すよ\n`!help`: このメッセージを表示するよ")
}
