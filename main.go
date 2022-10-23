package main

import (
	"flag"
	"log"
)

func main() {
	// event proccessor
	//consumer (fetcher, proccessor)
}

func mustToken() string {

	token := flag.String(
		"tg-bot-token",
		"",
		"token to access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
