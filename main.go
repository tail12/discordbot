package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	Token   string
	BotName string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Token = os.Getenv("DISCORD_API_TOKEN")
	BotName = os.Getenv("DISCORD_CLIENT_ID")
}

func main() {
	discord, err := discordgo.New()
	discord.Token = Token
	if err != nil {
		fmt.Println("Error logging")
		fmt.Println(err)
	}
	fmt.Println(Token)
	fmt.Println(BotName)
}
