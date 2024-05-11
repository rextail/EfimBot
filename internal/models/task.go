package models

type Task struct {
	ProjectCode     string `db:"project_code"`
	UserResponsible string `db:"user_responsible"`
	Description     string `db:"description"`
	DateGiven       string `db:"date_given"`
	Urgency         string `db:"urgency"`
}

type TaskView struct {
	PojectCode      string `db:"project_code"`
	Description     string `db:"description"`
	DateGiven       string `db:"date_given"`
	Urgency         string `db:"urgency"`
	UserResponsible int    `db:"user_id"`
}
