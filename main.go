package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ruribou/discord-bot-hinana/bot"
)

func main() {
	// ここで .env を読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_BOT_TOKEN is not set")
	}

	if err := bot.Start(token); err != nil {
		log.Fatalf("failed to start bot: %v", err)
	}
}
