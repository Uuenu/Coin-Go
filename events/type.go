package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Proccessor interface {
	Proccess(event Event) error
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type Type
	Text string
	Meta interface{}
}
