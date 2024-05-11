package events

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type  Type
	Value string
	Meta  interface{}
}

type Type int

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e []Event) error
}
