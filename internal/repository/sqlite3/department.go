package sqlite3

import (
	"EfimBot/internal/models"
	"EfimBot/pkg/e"
	"context"
	"database/sql"
)

type departmentRepo struct {
	departmentDB *sql.DB
}

func (d *departmentRepo) Create(ctx context.Context, department models.Department) error {
	query := `INSERT INTO deparments VALUES(?,?)`

	_, err := d.departmentDB.ExecContext(ctx, query, department.Name, department.Manager)
	if err != nil {
		return err
	}

	return nil
}

func (s *departmentRepo) Delete(ctx context.Context, name string) error {
	query := `DELETE FROM deparments WHERE department_name = ?`

	_, err := s.departmentDB.ExecContext(ctx, query, name)
	if err != nil {
		return err
	}

	return nil
}

func (s *departmentRepo) Get(ctx context.Context, name string) (deparment models.Department, err error) {
	query := `SELECT department_id FROM departments WHERE department_name = ?`

	err = s.departmentDB.QueryRowContext(ctx, query, name).Scan(&deparment.Name, &deparment.Manager)

	if err != nil {
		return models.Department{}, e.Wrap("can't get department", err)
	}

	return deparment, nil
}
