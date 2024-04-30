package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}
