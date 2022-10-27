package storage

import "time"

type Storage interface {
	AddRecord(p *Record) error
	RecordsList(chatID int, limit int) ([]Record, error)
	LastRecord(chatID int) (map[string]string, time.Time, error)
	UpdateLastRecord(chatID int, recData map[string]string) (err error)
}

type Record struct {
	Username string
	ChatID   int
	Time     time.Time
	Data     map[string]string
}
