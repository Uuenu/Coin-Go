package telegram

import (
	"telegram-coin-go/clients/telegram"
	"telegram-coin-go/storage"
)

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

type TgProcessor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}

func New(client *telegram.Client, storage storage.Storage) *TgProcessor {
	return &TgProcessor{
		tg:      client,
		storage: storage,
	}
}
