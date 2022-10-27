package telegram

import (
	"log"
	"strconv"
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

	log.Printf("got new command '%s' from '%s' ", text, username)

	m_float, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return p.saveRecord(chatID, m_float, username)
	}

	switch text {
	case AddRecord:
		//send messege "print sum"
		//p.sum(chatID, username)
	}

	return nil
}

func (p *TgProcessor) saveRecord(chatID int, sum float64, username string) (err error) {
	defer func() { err = lib.WrapIfErr("can't do command: save page", err) }()

	recData, recDate, err := events.RecData(chatID, sum, p.storage) // get data and time from last record

	timeNow := time.Now()

	if events.CheckTime(recDate, timeNow) {
		p.storage.UpdateLastRecord(chatID, recData)
	} else {
		record := &storage.Record{
			ChatID:   chatID,
			Username: username,
			Time:     time.Now(),
			Data:     recData,
		}
		if err := p.storage.AddRecord(record); err != nil {
			return err
		}
	}

	if err := p.tg.SendMessage(chatID, msgSaved); err != nil {
		return err
	}
	return nil
}
