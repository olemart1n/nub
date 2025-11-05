package utils

import (
	"fmt"
	"strings"
)

func BuildBulkInsertSQL(table string, columns []string, rows [][]any) (string, []any) {
	placeholders := make([]string, len(rows))
	values := make([]any, 0, len(rows)*len(columns))

	for i, row := range rows {
		ph := make([]string, len(columns))
		for j := range columns {
			ph[j] = fmt.Sprintf("$%d", i*len(columns)+j+1)
		}
		placeholders[i] = fmt.Sprintf("(%s)", strings.Join(ph, ", "))
		values = append(values, row...)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s",
		table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	return query, values
}
