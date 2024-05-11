package sqlite3

import (
	"EfimBot/storage"
	"fmt"
	"strings"
)

type InitHandler struct {
}

func (h InitHandler) GetInitQuery(tab storage.Table) string {
	typifiedColumns := h.getTypifiedColumnsWithPrimaryKey(tab)
	return fmt.Sprintf(`CREATE  IF NOT EXISTS %s (%s)`, tab.Name, typifiedColumns)
}

func (h InitHandler) getTypifiedColumnsWithPrimaryKey(tab storage.Table) string {

	//tab.Columns do not contain <tableName>_id to create primary key, then we just get it using table.Name
	//this also means that anything in columns that contains id is foreign key

	columns := strings.Split(tab.Columns, ",")

	res := make([]string, len(columns)+1)
	res[0] = fmt.Sprintf(`%s %s`, tab.Name+"_id", `INTEGER PRIMARY KEY AUTOINCREMENT`)

	for i := 0; i < len(columns); i++ {
		if strings.Contains(columns[i], "id") {
			res[i+1] = h.foreignKey(columns[i])
		} else {
			res[i+1] = fmt.Sprintf(`%s TEXT`, columns[i])
		}
	}

	return strings.Join(res, ",")
}

func (h InitHandler) foreignKey(column string) string {
	underscoreIndex := strings.IndexByte(column, '_')
	mainPart := fmt.Sprintf(`%s INTEGER`, column)
	foreignPart := fmt.Sprintf(`FOREIGN KEY(%s) REFERENCES %s(%s)`, column, column[:underscoreIndex], column)

	return fmt.Sprintf(`%s,%s`, mainPart, foreignPart)
}
