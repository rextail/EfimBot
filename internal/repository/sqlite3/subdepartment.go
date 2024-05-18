package sqlite3

import (
	"EfimBot/internal/models"
	"context"
	"database/sql"
)

type subDepartmentRepo struct {
	subDB *sql.DB
}

func (s *subDepartmentRepo) Get(ctx context.Context, name string) (sub models.SubDepartment, err error) {
	query := `SELECT subdeparment_id FROM subdepartments WHERE sub_name = ?`

	err = s.subDB.QueryRowContext(ctx, query, name).Scan(&sub.Name, &sub.Chief)
	if err != nil {
		return models.SubDepartment{}, err
	}

	return sub, nil
}

func (s *subDepartmentRepo) Create(ctx context.Context, name string, chief string) error {
	query := `INSERT INTO subdepartments VALUES(?,?)`

	_, err := s.subDB.ExecContext(ctx, query, name, chief)
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
