package telegram

import (
	"fmt"
	"log"
	"strconv"
	"telegram-coin-go/events"
	lib "telegram-coin-go/lib/e"
	"telegram-coin-go/storage"
	"time"
)

const (
	RndCmd       = "/rnd"
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

	// switch text {
	// case AddRecord:
	// 	//send messege "print sum"
	// 	//p.sum(chatID, username)
	// }

	return nil
}

func (p *TgProcessor) saveRecord(chatID int, sum float64, username string) (err error) {
	defer func() { err = lib.WrapIfErr("can't do command: save page", err) }()

	_, recDate, err := events.LastData(chatID, sum, p.storage) // get data and time from last record
	timeNow := time.Now().Format("2006-January-02")

	if timeNow == recDate {
		fmt.Println("TIME THE SAME")
		p.storage.UpdateLastRecord(chatID, sum)
	} else {
		record := &storage.Record{
			ChatID:   chatID,
			Username: username,
			Time:     time.Now().Format("2006-January-02"),
			Data:     events.NewData(sum),
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
