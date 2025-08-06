package dialects

import (
	"fmt"
	"strings"
)

type GenericDialect struct {
	BaseDialect
}

func (d *GenericDialect) Name() string {
	return "generic"
}

func (d *GenericDialect) QuoteIdentifier(identifier string) string {
	return fmt.Sprintf("\"%s\"", identifier)
}

func (d *GenericDialect) CreateTable(tableName string, columns []ColumnDef) string {
	var sb strings.Builder

	sb.WriteString("CREATE TABLE ")
	sb.WriteString(d.QuoteIdentifier(tableName))
	sb.WriteString(" (\n")

	// First add all columns
	for i, col := range columns {
		if i > 0 {
			sb.WriteString(",\n")
		}

		sb.WriteString("  ")
		sb.WriteString(d.QuoteIdentifier(col.Name))
		sb.WriteString(" ")
		sb.WriteString(string(col.Type))

		if col.IsPrimaryKey {
			sb.WriteString(" PRIMARY KEY")
		} else if !col.Nullable {
			sb.WriteString(" NOT NULL")
		}
	}

	// Then add foreign key constraints
	for _, col := range columns {
		if col.IsForeignKey && col.References != "" {
			// Parse the reference (table.column)
			parts := strings.Split(col.References, ".")
			if len(parts) != 2 {
				continue
			}
			refTable := parts[0]
			refColumn := parts[1]

			sb.WriteString(",\n  FOREIGN KEY (")
			sb.WriteString(d.QuoteIdentifier(col.Name))
			sb.WriteString(") REFERENCES ")
			sb.WriteString(d.QuoteIdentifier(refTable))
			sb.WriteString(" (")
			sb.WriteString(d.QuoteIdentifier(refColumn))
			sb.WriteString(")")
		}
	}

	sb.WriteString("\n);\n")

	return sb.String()
}

func (d *GenericDialect) InsertInto(tableName string, columns []string, values [][]interface{}, batchSize int) string {
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
