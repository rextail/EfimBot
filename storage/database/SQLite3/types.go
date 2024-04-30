package sqlite3

import (
	"fmt"
	"strings"
)

type Table struct {
	Name      string
	Columns   string
	Dimension int
}

func (table Table) GetInsertQuery() string {
	values := fmt.Sprintf("VALUES (%s)", table.valuesSequence())
	columns := fmt.Sprintf("(%s)", table.Columns)
	return fmt.Sprintf(`INSERT INTO %s %s %s`, table.Name, columns, values)
}

func (table Table) valuesSequence() string {
	return strings.Join(strings.Split(strings.Repeat("?", table.Dimension), ""), ",")
}

func (table Table) GetInitQuery() string {
	columns := table.getTypifiedColumns()

	return fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s)`, table.Name, columns)
}

func (table Table) getTypifiedColumns() string {
	var underscoreIndex int
	columns := strings.Split(table.Columns, ",")

	columns[0] = fmt.Sprintf(`%s %s`, columns[0], `INTEGER PRIMARY KEY AUTOINCREMENT`)

	for i := 1; i < len(columns); i++ {
		if strings.Contains(columns[i], "id") {
			underscoreIndex = strings.IndexByte(columns[i], '_')
			columns[i] = fmt.Sprintf(`FOREIGN KEY(%s) REFERENCES %s(%s)`, columns[i], columns[i][:underscoreIndex], columns[i])
		} else {
			columns[i] = fmt.Sprintf(`%s TEXT`, columns[i])
		}
	}
	return strings.Join(columns, ",")
}
