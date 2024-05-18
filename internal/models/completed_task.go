package models

type completedTask struct {
	ID                int    `db:"task_id"`
	WorkerResponsible int    `db:"worker_id"`
	ProjectCode       string `db:"project_code"`
	Description       string `db:"description"`
	Type              string `db:"type"`
	CompletionDate    string `db:"completion_date"`
}
