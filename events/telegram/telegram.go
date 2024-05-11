package telegram

import (
	"EfimBot/clients/telegram"
	"EfimBot/e"
	"EfimBot/events"
	"EfimBot/storage"
	"errors"
)

var ErrUnknownEvent = errors.New("Unknown event type")
var ErrUnknownMetaType = errors.New("Unknown meta type")

type Processor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}

type Meta struct {
	ChatID   int
	Username string
}

func (p *Processor) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		p.processMessage(event)
	default:
		return e.Wrap("error occured during message processing", ErrUnknownEvent)
	}
	return nil
}

func (p *Processor) processMessage(event events.Event) error {
	meta, err := meta(event)
	if err != nil {
		return e.Wrap("error occured during message processing", err)
	}
}

func meta(event events.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)

	if !ok {
		return Meta{}, ErrUnknownMetaType
	}

	return res, nil
}

func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, e.Wrap("Can't get events", err)
	}

	res := make([]events.Event, 0, len(updates))

	for _, u := range updates {
		res = append(res, event(u))
	}

	p.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

func event(upd telegram.Update) events.Event {
	updType := fetchType(upd)

	res := events.Event{
		Type:  updType,
		Value: fetchText(upd),
	}

	if updType == events.Message {
		res.Meta = Meta{
			ChatID:   upd.Message.Chat.ID,
			Username: upd.Message.From.Username,
		}
	}

	return res
}

func fetchText(upd telegram.Update) string {
	if upd.Message == nil {
		return ""
	}
	return upd.Message.Text
}

func fetchType(upd telegram.Update) events.Type {
	if upd.Message == nil {
		return events.Unknown
	}
	return events.Message
}
