package dialects

import (
	"fmt"
	"strings"
)

type OracleDialect struct {
	BaseDialect
}

func (d *OracleDialect) Name() string {
	return "oracle"
}

func (d *OracleDialect) QuoteIdentifier(identifier string) string {
	return fmt.Sprintf("\"%s\"", strings.ToUpper(identifier))
}

func (d *OracleDialect) CreateTable(tableName string, columns []ColumnDef) string {
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

		oracleType := d.mapSQLType(col.Type)
		sb.WriteString(oracleType)

		if !col.Nullable {
			sb.WriteString(" NOT NULL")
		}
	}

	sb.WriteString("\n);\n")

	return sb.String()
}

func (d *OracleDialect) InsertInto(tableName string, columns []string, values [][]interface{}, batchSize int) string {
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

		if len(batch) > 1 {

			sb.WriteString("INSERT ALL\n")

			for _, row := range batch {
				sb.WriteString("  INTO ")
				sb.WriteString(d.QuoteIdentifier(tableName))
				sb.WriteString(" (")

				for i, col := range columns {
					if i > 0 {
						sb.WriteString(", ")
					}
					sb.WriteString(d.QuoteIdentifier(col))
				}

				sb.WriteString(") VALUES (")

				for i, val := range row {
					if i > 0 {
						sb.WriteString(", ")
					}
					sb.WriteString(d.FormatValue(val))
				}

				sb.WriteString(")\n")
			}

			sb.WriteString("SELECT 1 FROM DUAL;\n\n")
		} else if len(batch) == 1 {

			row := batch[0]

			sb.WriteString("INSERT INTO ")
			sb.WriteString(d.QuoteIdentifier(tableName))
			sb.WriteString(" (")

			for i, col := range columns {
				if i > 0 {
					sb.WriteString(", ")
				}
				sb.WriteString(d.QuoteIdentifier(col))
			}

			sb.WriteString(") VALUES (")

			for i, val := range row {
				if i > 0 {
					sb.WriteString(", ")
				}
				sb.WriteString(d.FormatValue(val))
			}

			sb.WriteString(");\n\n")
		}
	}

	return sb.String()
}

func (d *OracleDialect) mapSQLType(sqlType SQLType) string {
	switch sqlType {
	case SQLTypeInteger:
		return "NUMBER(10)"
	case SQLTypeFloat:
		return "NUMBER"
	case SQLTypeText:
		return "CLOB"
	case SQLTypeDate:
		return "DATE"
	case SQLTypeDateTime:
		return "TIMESTAMP"
	case SQLTypeBoolean:
		return "NUMBER(1)" // Oracle uses 0 and 1 for booleans
	default:
		return string(sqlType)
	}
}
