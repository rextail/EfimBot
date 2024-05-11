package models

type Project struct {
	Code           string `db:"code_id"`
	Name           string `db:"project_name"`
	ProjectManager string `db:"project_manager"`
}
