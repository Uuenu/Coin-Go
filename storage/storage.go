package storage

import "time"

type Storage interface {
	AddRecord(p *Page) error
	RecordsList(limit int) ([]Page, error)
}

type Page struct {
	Data     map[string]string
	Username string
	ChatID   int
	Date     time.Time
}
