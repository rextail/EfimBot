package models

type Task struct {
	ID                int    `db:"task_id"`
	WorkerResponsible int    `db:"worker_id"`
	ProjectCode       string `db:"project_code"`
	Description       string `db:"description"`
	Type              string `db:"type"`
	DateGiven         string `db:"date_given"`
}
