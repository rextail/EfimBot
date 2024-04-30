package sqlite3

import (
	"EfimBot/storage"
	"context"
	"database/sql"
	"fmt"
	"strings"
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
	if err != nil{
		return nil, err
	}

	return result, nil
}

func (s *Storage) Insert(ctx context.Context, table storage.TableInfo, parameters []string) error {
	query := formInsertQuery(table)

	_, err := s.db.ExecContext(ctx, query, parameters)

	if err != nil {
		return err
	}

	return nil
}

func formInsertQuery(table storage.TableInfo) string {
	values := fmt.Sprintf("VALUES (%s)", formValuesSequence(table.Type))
	columns := fmt.Sprintf("(%s)", table.Columns)
	return fmt.Sprintf(`INSERT INTO %s %s %s`, table.Name, columns, values)
}

func formValuesSequence(length int) string {
	return strings.Join(strings.Split(strings.Repeat("?", length), ""), ",")
}

func (s *Storage) Init(ctx context.Context, table storage.TableInfo) error {
	query := formInitQuery(table)
}

func formInitQuery(table storage.TableInfo) string {
	columns := 

	return fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s)`)
}