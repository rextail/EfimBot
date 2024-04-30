package telegram

import (
	"EfimBot/clients/telegram"
	"EfimBot/storage"
)

type Processor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}
