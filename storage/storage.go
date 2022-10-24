package storage

import "time"

type Storage interface {
	AddRecord(p *Page) error
	RecordsList() []Page
}

type Page struct {
	Data     map[string]string
	Username string
	ChatID   int
	Date     time.Time
}
