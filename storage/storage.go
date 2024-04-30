package storage

import (
	"context"
	"database/sql"
)

type Storage interface {
	Insert(ctx context.Context, table Table, parameters []string) error
	Select(ctx context.Context, query string, parameters []string) (*sql.Rows, error)
}

type Table interface {
	GetInsertQuery() string
	GetInitQuery() string
}
