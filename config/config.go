package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadToken() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_BOT_TOKEN is not set")
	}

	return token
}
