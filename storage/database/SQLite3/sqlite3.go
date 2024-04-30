package sqlite3

import (
	"context"
	"database/sql"
)

type Storage struct {
	db *sql.DB
}

func New(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Select(ctx context.Context, query string, parameters []string) (*sql.Rows, error) {
	result, err := s.db.QueryContext(ctx, query, parameters)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Storage) Insert(ctx context.Context, table Table, parameters []string) error {
	_, err := s.db.ExecContext(ctx, table.GetInsertQuery(), parameters)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Init(ctx context.Context, table Table) error {
	_, err := s.db.ExecContext(ctx, table.GetInitQuery())
	if err != nil {
		return err
	}

	return nil
}
