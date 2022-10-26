package storage

import "time"

type Storage interface {
	AddRecord(p *Record) error
	RecordsList(limit int) ([]Record, error)
	LastCredit(chatID int) (string, string, error)
}

type Record struct {
	Username string
	ChatID   int
	Time     time.Time
	Data     map[string]string
}
