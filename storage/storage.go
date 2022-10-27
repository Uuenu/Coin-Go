package storage

type Storage interface {
	AddRecord(p *Record) error
	RecordsList(chatID int, limit int) ([]Record, error)
	LastRecord(chatID int) (map[string]string, string, error)
	UpdateLastRecord(chatID int, sum float64) (err error)
}

type Record struct {
	Username string
	ChatID   int
	Time     string
	Data     map[string]string
}
