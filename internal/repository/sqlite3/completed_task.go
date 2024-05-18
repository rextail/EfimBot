package sqlite3

import (
	"EfimBot/internal/models"
	"EfimBot/pkg/e"
	"context"
	"database/sql"
)

type completedTaskRepo struct {
	completedRepo *sql.DB
}

func (c *completedTaskRepo) Create(ctx context.Context, task models.Task) error {
	query := `INSERT INTO completed_tasks VALUES(?,?,?,?,?,?)`

	_, err := c.completedRepo.ExecContext(ctx, query, task.ProjectCode, task.Description, task.WorkerResponsible, task.Type, taskTime())
	if err != nil {
		return e.Wrap("can't make task completed", err)
	}

	return nil
}
