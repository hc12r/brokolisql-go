package dialects

import (
	"fmt"
	"strings"
)

type SQLiteDialect struct {
	BaseDialect
}

func (d *SQLiteDialect) Name() string {
	return "sqlite"
}

func (d *SQLiteDialect) QuoteIdentifier(identifier string) string {
	return fmt.Sprintf("\"%s\"", identifier)
}

func (d *SQLiteDialect) CreateTable(tableName string, columns []ColumnDef) string {
	var sb strings.Builder

	sb.WriteString("CREATE TABLE ")
	sb.WriteString(d.QuoteIdentifier(tableName))
	sb.WriteString(" (\n")

	for i, col := range columns {
		if i > 0 {
			sb.WriteString(",\n")
		}

		sb.WriteString("  ")
		sb.WriteString(d.QuoteIdentifier(col.Name))
		sb.WriteString(" ")

		sqliteType := d.mapSQLType(col.Type)
		sb.WriteString(sqliteType)

		if !col.Nullable {
			sb.WriteString(" NOT NULL")
		}
	}

	sb.WriteString("\n);\n")

	return sb.String()
}

func (d *SQLiteDialect) InsertInto(tableName string, columns []string, values [][]interface{}, batchSize int) string {
	var sb strings.Builder

	if batchSize <= 0 {
		batchSize = len(values)
	}

	for batchStart := 0; batchStart < len(values); batchStart += batchSize {
		batchEnd := batchStart + batchSize
		if batchEnd > len(values) {
			batchEnd = len(values)
		}

		batch := values[batchStart:batchEnd]

		sb.WriteString("INSERT INTO ")
		sb.WriteString(d.QuoteIdentifier(tableName))
		sb.WriteString(" (")

		for i, col := range columns {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(d.QuoteIdentifier(col))
		}

		sb.WriteString(") VALUES\n")

		for i, row := range batch {
			if i > 0 {
				sb.WriteString(",\n")
			}

			sb.WriteString("(")

			for j, val := range row {
				if j > 0 {
					sb.WriteString(", ")
				}
				sb.WriteString(d.FormatValue(val))
			}

			sb.WriteString(")")
		}

		sb.WriteString(";\n\n")
	}

	return sb.String()
}

func (d *SQLiteDialect) mapSQLType(sqlType SQLType) string {
	switch sqlType {
	case SQLTypeInteger:
		return "INTEGER"
	case SQLTypeFloat:
		return "REAL"
	case SQLTypeText:
		return "TEXT"
	case SQLTypeDate, SQLTypeDateTime:
		return "TEXT" // SQLite doesn't have native date types
	case SQLTypeBoolean:
		return "INTEGER" // SQLite uses 0 and 1 for booleans
	default:
		return string(sqlType)
	}
}
