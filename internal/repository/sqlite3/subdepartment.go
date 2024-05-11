package sqlite3

import (
	"context"
	"database/sql"
)

type subDepartmentRepo struct {
	subDB *sql.DB
}

func (s *subDepartmentRepo) GetID(ctx context.Context, name string) (ID int, err error) {
	query := `SELECT sub_name FROM subdepartments WHERE sub_name = ?`

	err = s.subDB.QueryRowContext(ctx, query, name).Scan(&ID)
	if err != nil {
		return -1, err
	}

	return ID, nil
}

func (s *subDepartmentRepo) Create(ctx context.Context, name string, chief string) error {
	query := `INSERT INTO subdepartments VALUES (?,?)`

	args := []string{name, chief}

	_, err := s.subDB.ExecContext(ctx, query, args)
	if err != nil {
		return err
	}

	return nil
}

func (s *subDepartmentRepo) Delete(ctx context.Context, name string) error {
	query := `DELETE FROM subdepartments WHERE sub_name = ?`

	_, err := s.subDB.ExecContext(ctx, query, name)
	if err != nil {
		return err
	}

	return nil
}
