package dto

type Task struct {
	ProjectCode string `db:"project_code"`
	Description string `db:"description"`
}
