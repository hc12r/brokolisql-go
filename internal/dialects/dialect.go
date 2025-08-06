package dialects

import (
	"fmt"
	"strings"
)

type SQLType string

const (
	SQLTypeInteger  SQLType = "INTEGER"
	SQLTypeFloat    SQLType = "FLOAT"
	SQLTypeText     SQLType = "TEXT"
	SQLTypeDate     SQLType = "DATE"
	SQLTypeDateTime SQLType = "DATETIME"
	SQLTypeBoolean  SQLType = "BOOLEAN"
)

type ColumnDef struct {
	Name         string
	Type         SQLType
	Nullable     bool
	IsPrimaryKey bool
	IsForeignKey bool
	References   string // Referenced table.column
}

type Dialect interface {
	Name() string

	CreateTable(tableName string, columns []ColumnDef) string

	InsertInto(tableName string, columns []string, values [][]interface{}, batchSize int) string

	QuoteIdentifier(identifier string) string

	FormatValue(value interface{}) string
}

func GetDialect(name string) (Dialect, error) {
	name = strings.ToLower(name)

	switch name {
	case "postgres", "postgresql":
		return &PostgresDialect{}, nil
	case "mysql":
		return &MySQLDialect{}, nil
	case "sqlite":
		return &SQLiteDialect{}, nil
	case "sqlserver", "mssql":
		return &SQLServerDialect{}, nil
	case "oracle":
		return &OracleDialect{}, nil
	case "generic":
		return &GenericDialect{}, nil
	default:
		return nil, fmt.Errorf("unsupported SQL dialect: %s", name)
	}
}

type BaseDialect struct{}

func (d *BaseDialect) FormatValue(value interface{}) string {
	if value == nil {
		return "NULL"
	}

	switch v := value.(type) {
	case string:

		escaped := strings.ReplaceAll(v, "'", "''")
		return fmt.Sprintf("'%s'", escaped)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v)
	case float32, float64:
		return fmt.Sprintf("%g", v)
	case bool:
		if v {
			return "TRUE"
		}
		return "FALSE"
	default:

		str := fmt.Sprintf("%v", v)
		escaped := strings.ReplaceAll(str, "'", "''")
		return fmt.Sprintf("'%s'", escaped)
	}
}
