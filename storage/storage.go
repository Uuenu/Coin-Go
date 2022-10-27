package storage

import "time"

type Storage interface {
	AddRecord(p *Record) error
	RecordsList(limit int) ([]Record, error)
	LastRecord(chatID int) (string, string, time.Time, error)
	UpdateLastRecord(chatID int, recData map[string]string) (err error)
}

type Record struct {
	Username string
	ChatID   int
	Time     time.Time
	Data     map[string]string
}
