package telegram

import (
	"fmt"
	"log"
	"strconv"
	lib "telegram-coin-go/lib/e"
	"telegram-coin-go/storage"
	"time"
)

const (
	Today        = "/today"
	HelpCmd      = "/help"
	StartCmd     = "/start"
	AddRecord    = "/record"
	DebitnCredit = "/upd"
)

func (p *TgProcessor) doCmd(text string, chatID int, username string) error {

	log.Printf("got new command '%s' from '%s' ", text, username)

	m_float, err := strconv.ParseFloat(text, 64)
	fmt.Println(m_float, " ", err)
	if err == nil && m_float != 0 {
		fmt.Println("here")
		return p.saveRecord(chatID, m_float, username)
	}

	switch text {
	case Today:
		p.Today(chatID)
	}

	return nil
}

func (p *TgProcessor) Today(chatID int) error {
	result, err := p.storage.Today(chatID)
	if err != nil {
		return err
	}
	p.tg.SendMessage(chatID, "Total: "+strconv.FormatFloat(result.Day.Total, 'f', 5, 64))
	return nil
}

func (p *TgProcessor) saveRecord(chatID int, sum float64, username string) (err error) {
	defer func() { err = lib.WrapIfErr("can't do command: save page", err) }()

	timeNow := time.Now().Format("2006-January-02")

	if p.storage.CheckTime(chatID, timeNow) {
		p.storage.UpdateLastRecord(chatID, sum)
	} else {
		record := &storage.Record{
			ChatID:   chatID,
			Username: username,
			Time:     time.Now().Format("2006-January-02"),
			Day: storage.Day{
				Total: sum,
			},
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
