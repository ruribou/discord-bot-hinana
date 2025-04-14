package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func HelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if _, err := s.ChannelMessageSend(m.ChannelID, "`!ping`: Pong! を返すよ\n`!help`: このメッセージを表示するよ"); err != nil {
		// Log the error or handle it as needed
		fmt.Println("Error sending message:", err)
	}
}
