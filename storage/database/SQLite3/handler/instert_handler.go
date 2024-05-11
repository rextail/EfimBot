package insert

import (
	"EfimBot/storage"
	"fmt"
	"strings"
)

type InsertHandler struct {
}

func (h InsertHandler) GetInsertQuery(tab storage.Table) string {
	values := fmt.Sprintf("VALUES (%s)", h.ValuesSequence(tab.Dimension))
	columns := fmt.Sprintf("(%s)", tab.Columns)
	return fmt.Sprintf(`INSERT INTO %s %s %s`, tab.Name, columns, values)
}

func (h InsertHandler) ValuesSequence(dimension int) string {
	return strings.Join(strings.Split(strings.Repeat("?", dimension), ""), ",")
}
