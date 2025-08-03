package dialects

import (
	"fmt"
	"strings"
)

type MySQLDialect struct {
	BaseDialect
}

func (d *MySQLDialect) Name() string {
	return "mysql"
}

func (d *MySQLDialect) QuoteIdentifier(identifier string) string {
	return fmt.Sprintf("`%s`", identifier)
}

func (d *MySQLDialect) CreateTable(tableName string, columns []ColumnDef) string {
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

		mysqlType := d.mapSQLType(col.Type)
		sb.WriteString(mysqlType)

		if !col.Nullable {
			sb.WriteString(" NOT NULL")
		}
	}

	sb.WriteString("\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;\n")

	return sb.String()
}

func (d *MySQLDialect) InsertInto(tableName string, columns []string, values [][]interface{}, batchSize int) string {
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

func (d *MySQLDialect) mapSQLType(sqlType SQLType) string {
	switch sqlType {
	case SQLTypeInteger:
		return "INT"
	case SQLTypeFloat:
		return "DOUBLE"
	case SQLTypeText:
		return "TEXT"
	case SQLTypeDate:
		return "DATE"
	case SQLTypeDateTime:
		return "DATETIME"
	case SQLTypeBoolean:
		return "TINYINT(1)"
	default:
		return string(sqlType)
	}
}
