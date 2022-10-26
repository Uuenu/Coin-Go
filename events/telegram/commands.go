package telegram

import (
	"telegram-coin-go/events"
	lib "telegram-coin-go/lib/e"
	"telegram-coin-go/storage"
	"time"
)

const (
	RndCmd    = "/rnd"
	HelpCmd   = "/help"
	StartCmd  = "/start"
	AddRecord = "/record"
)

func (p *TgProcessor) doCmd(text string, chatID int, username string) error {

	switch text {
	case AddRecord:
		//send messege "print sum"

	}

	return nil
}

func (p *TgProcessor) saveRecord(chatID int, sum string, username string) (err error) {
	defer func() { err = lib.WrapIfErr("can't do command: save page", err) }()

	data, err := events.RecData(chatID, sum, p.storage)

	record := &storage.Record{
		ChatID:   chatID,
		Username: username,
		Time:     time.Now(),
		Data:     data,
	}

	if err := p.storage.AddRecord(record); err != nil {
		return err
	}

	if err := p.tg.SendMessage(chatID, msgSaved); err != nil {
		return err
	}
	return nil
}
