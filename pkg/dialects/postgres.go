package dialects

import (
	"fmt"
	"strings"
)

type PostgresDialect struct {
	BaseDialect
}

func (d *PostgresDialect) Name() string {
	return "postgresql"
}

func (d *PostgresDialect) QuoteIdentifier(identifier string) string {
	return fmt.Sprintf("\"%s\"", identifier)
}

func (d *PostgresDialect) CreateTable(tableName string, columns []ColumnDef) string {
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

		pgType := d.mapSQLType(col.Type)
		sb.WriteString(pgType)

		if !col.Nullable {
			sb.WriteString(" NOT NULL")
		}
	}

	sb.WriteString("\n);\n")

	return sb.String()
}

func (d *PostgresDialect) InsertInto(tableName string, columns []string, values [][]interface{}, batchSize int) string {
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

func (d *PostgresDialect) mapSQLType(sqlType SQLType) string {
	switch sqlType {
	case SQLTypeInteger:
		return "INTEGER"
	case SQLTypeFloat:
		return "DOUBLE PRECISION"
	case SQLTypeText:
		return "TEXT"
	case SQLTypeDate:
		return "DATE"
	case SQLTypeDateTime:
		return "TIMESTAMP"
	case SQLTypeBoolean:
		return "BOOLEAN"
	default:
		return string(sqlType)
	}
}
