package main

import (
	"flag"
	"log"
)

func main() {
	const host = "api.telegram.org"
}

func mustToken() string {
	token := flag.String("Efim-bot-token", "", "Token to authorize the bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("Bot token is not specified")
	}

	return *token
}
