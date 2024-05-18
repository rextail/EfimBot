package sqlite3

import (
	"EfimBot/internal/models"
	"EfimBot/pkg/e"
	"context"
	"database/sql"
	"time"
)

type taskRepo struct {
	taskDB *sql.DB
}

func (t *taskRepo) CreateByCodeAndDescription(ctx context.Context, code string, description string) error {
	query := `INSERT INTO task(project_code,description) VALUES(?,?)`

	_, err := t.taskDB.ExecContext(ctx, query, code, description)
	if err != nil {
		return err
	}

	return nil
}

func (t *taskRepo) WorkerTasks(ctx context.Context, worker models.Worker) (tasks []models.Task, err error) {
	query := `SELECT task_id, description FROM tasks WHERE worker_id = ?`

	rows, err := t.taskDB.QueryContext(ctx, query, worker.ID)
	if err != nil {
		return nil, e.Wrap("can't execute worker's tasks", err)
	}
	for rows.Next() {
		var task models.Task
		rows.Scan(&task.ID, &task.WorkerResponsible, &task.ProjectCode, &task.Description, &task.Type, &task.DateGiven)
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t *taskRepo) AdressInternalTo(ctx context.Context, task_id int, worker models.Worker) error {
	query := `UPDATE tasks SET worker_responsible = ?, type = ?, date_given = ? WHERE task_id = ?`

	_, err := t.taskDB.ExecContext(ctx, query, worker.Name, "Внутренняя", taskTime(), task_id)
	if err != nil {
		return err
	}

	return nil
}

func (t *taskRepo) AdressExternalTo(ctx context.Context, task_id int, department models.Department) error {
	query := `UPDATE tasks SET worker_responsible = ?, type = ?, date_given = ?, WHERE task_id = ?`

	_, err := t.taskDB.ExecContext(ctx, query, department.Manager, "Внешняя", taskTime(), task_id)
	if err != nil {
		return err
	}

	return nil
}

func (t *taskRepo) Exist(ctx context.Context, task_id int) (bool, error) {
	query := `SELECT * FROM tasks WHERE task_id = ?`

	res, err := t.taskDB.ExecContext(ctx, query, task_id)
	if err != nil {
		return false, e.Wrap("can't process exists command", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return false, e.Wrap("driver may not support counting affected rows", err)
	}
	if affected == 0 {
		return false, nil
	}

	return true, nil
}

func taskTime() string {
	return time.Now().Format("01/02/06")
}
