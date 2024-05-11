package models

type User struct {
	Name          string `db:"name"`
	Position      string `db:"position"`
	Department    string `db:"department"`
	SubDepartment string `db:"sub_department"`
}
