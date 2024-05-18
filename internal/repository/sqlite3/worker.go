package sqlite3

import (
	"EfimBot/internal/models"
	"EfimBot/pkg/e"
	"context"
	"database/sql"
)

type workerRepo struct {
	workerDB *sql.DB
}

func (u *workerRepo) Create(ctx context.Context, worker models.Worker) error {
	query := `INSERT INTO workers VALUES(?,?,?,?)`

	_, err := u.workerDB.ExecContext(ctx, query, worker.Name, worker.Position, worker.Department, worker.SubDepartment)
	if err != nil {
		return err
	}

	return nil
}

func (u *workerRepo) Delete(ctx context.Context, name string, subdepartment string) error {
	query := `DELETE FROM workers WHERE worker_name = ? AND sub_department = ?`

	_, err := u.workerDB.ExecContext(ctx, query, name, subdepartment)
	if err != nil {
		return err
	}

	return nil
}

func (u *workerRepo) GetID(ctx context.Context, name string, subdepartment string) (id int, err error) {
	query := `SELECT worker_id FROM users WHERE name = ? AND sub_department = ?`

	err = u.workerDB.QueryRowContext(ctx, query, name, subdepartment).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (u *workerRepo) Get(ctx context.Context, workerName string, subdepartment string) (worker models.Worker, err error) {
	query := `SELECT worker_id FROM workers WHERE worker_name = ? AND sub_department = ?`

	err = u.workerDB.QueryRowContext(ctx, query, workerName, subdepartment).Scan(&worker.ID,
		&worker.Name,
		&worker.Position,
		&worker.Department,
		&worker.SubDepartment)
	if err != nil {
		return models.Worker{}, e.Wrap("error occured, can't get worker from db", err)
	}

	return worker, nil
}

func (u *workerRepo) AllDepartmentWorkers(ctx context.Context, department string) (workers []models.Worker, err error) {
	query := `SELECT name, position, department, subdepartment FROM workers WHERE department = ?`

	rows, err := u.workerDB.QueryContext(ctx, query, department)

	for rows.Next() {
		var worker models.Worker
		if err = rows.Scan(
			&worker.ID,
			&worker.Name,
			&worker.Position,
			&worker.Department,
			&worker.SubDepartment,
		); err != nil {
			return nil, e.Wrap("can't execute department workers", err)
		}
		workers = append(workers, worker)
	}
	return workers, nil
}

func (u *workerRepo) AllSubDepartmentWorkers(ctx context.Context, subdepartment string) (workers []models.Worker, err error) {
	query := `SELECT name,position, department, subdepartment FROM workers WHERE sub_department = ?`

	rows, err := u.workerDB.QueryContext(ctx, query, subdepartment)

	for rows.Next() {
		var worker models.Worker
		if err = rows.Scan(
			&worker.ID,
			&worker.Name,
			&worker.Position,
			&worker.Department,
			&worker.SubDepartment,
		); err != nil {
			return nil, e.Wrap("can't execute department workers", err)
		}
		workers = append(workers, worker)
	}
	return workers, nil
}
