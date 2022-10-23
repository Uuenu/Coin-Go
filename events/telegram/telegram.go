package telegram

import (
	"telegram-coin-go\clients\telegram"
	"telegram-coin-go\storage"
)

type TgProcessor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}

func New(client *telegra.Client, storage storage.Storage) *TgProcessor {
	return &TgProcessor{
		tg: client, 
		storage: storage,
	}	
}