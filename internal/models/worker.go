package models

type Worker struct {
	ID            int    `db:"worker_id"`
	Name          string `db:"name"`
	Position      string `db:"position"`
	Department    string `db:"department"`
	SubDepartment string `db:"sub_department"`
}
