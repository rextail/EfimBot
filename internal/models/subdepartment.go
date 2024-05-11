package models

type SubDepartment struct {
	Name  string `db:"sub_name"`
	Chief string `db:"chief_name"`
}
