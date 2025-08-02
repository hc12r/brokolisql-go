package dialects

import (
	"fmt"
	"strings"
)

type SQLServerDialect struct {
	BaseDialect
}

func (d *SQLServerDialect) Name() string {
	return "sqlserver"
}

func (d *SQLServerDialect) QuoteIdentifier(identifier string) string {
	return fmt.Sprintf("[%s]", identifier)
}

func (d *SQLServerDialect) CreateTable(tableName string, columns []ColumnDef) string {
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

		sqlServerType := d.mapSQLType(col.Type)
		sb.WriteString(sqlServerType)

		if !col.Nullable {
			sb.WriteString(" NOT NULL")
		}
	}

	sb.WriteString("\n);\n")

	return sb.String()
}

func (d *SQLServerDialect) InsertInto(tableName string, columns []string, values [][]interface{}, batchSize int) string {
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

		sb.WriteString(")\n")

		if len(batch) > 1 {
			sb.WriteString("SELECT * FROM (\n")
		} else {
			sb.WriteString("VALUES\n")
		}

		for i, row := range batch {
			if i > 0 {
				sb.WriteString(" UNION ALL\n")
			}

			if len(batch) > 1 {
				sb.WriteString("  SELECT ")
				for j, val := range row {
					if j > 0 {
						sb.WriteString(", ")
					}
					sb.WriteString(d.FormatValue(val))
				}
			} else {
				sb.WriteString("(")
				for j, val := range row {
					if j > 0 {
						sb.WriteString(", ")
					}
					sb.WriteString(d.FormatValue(val))
				}
				sb.WriteString(")")
			}
		}

		if len(batch) > 1 {
			sb.WriteString("\n) AS source")
		}

		sb.WriteString(";\n\n")
	}

	return sb.String()
}

func (d *SQLServerDialect) mapSQLType(sqlType SQLType) string {
	switch sqlType {
	case SQLTypeInteger:
		return "INT"
	case SQLTypeFloat:
		return "FLOAT"
	case SQLTypeText:
		return "NVARCHAR(MAX)"
	case SQLTypeDate:
		return "DATE"
	case SQLTypeDateTime:
		return "DATETIME2"
	case SQLTypeBoolean:
		return "BIT"
	default:
		return string(sqlType)
	}
}
