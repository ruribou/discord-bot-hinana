package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/ruribou/discord-bot-hinana/bot/commands"
)

func Start(token string) error {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return fmt.Errorf("error creating Discord session: %w", err)
	}

	dg.Identify.Intents = discordgo.IntentsGuildMessages
	dg.AddHandler(commands.MessageCreateHandler)

	if err := dg.Open(); err != nil {
		return fmt.Errorf("error opening connection: %w", err)
	}
	defer dg.Close()

	fmt.Println("Bot is running. Press Ctrl+C to exit.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	return nil
}
