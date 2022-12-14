package main

import (
	"flag"
	"log"

	tgClient "telegram-coin-go/clients/telegram"
	eventconsumer "telegram-coin-go/consumer/event-consumer"
	"telegram-coin-go/events/telegram"

	"telegram-coin-go/storage/mongodb"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {
	eventsProccessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		mongodb.New(), // monogdb

	)

	log.Print("service started")

	consumer := eventconsumer.New(eventsProccessor, eventsProccessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal()
	}

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
