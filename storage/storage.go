package storage

type Storage interface {
	AddRecord(p *Record) error
	DaysList(chatID int, limit int) ([]Record, error)
	LastRecord(chatID int) (result Record, err error)
	UpdateLastRecord(chatID int, sum float64) (err error)
	CheckTime(ChatID int, TimeNow string) bool
}

type Record struct {
	ChatID   int    `json:"chat_id" bson:"chat_id"`
	Username string `json:"username" bson:"username"`
	Time     string `json:"time" bson:"time"`
	Day      Day    `json:"day" bson:"day"`
}

type Day struct {
	Total float64 `json:"total" bson:"total"`
}
