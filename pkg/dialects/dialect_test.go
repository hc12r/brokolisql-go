package dialects

import (
	"testing"
)

func TestGetDialect(t *testing.T) {
	tests := []struct {
		name     string
		dialect  string
		wantType interface{}
		wantErr  bool
	}{
		{
			name:     "PostgreSQL",
			dialect:  "postgres",
			wantType: &PostgresDialect{},
			wantErr:  false,
		},
		{
			name:     "PostgreSQL (alternative name)",
			dialect:  "postgresql",
			wantType: &PostgresDialect{},
			wantErr:  false,
		},
		{
			name:     "MySQL",
			dialect:  "mysql",
			wantType: &MySQLDialect{},
			wantErr:  false,
		},
		{
			name:     "SQLite",
			dialect:  "sqlite",
			wantType: &SQLiteDialect{},
			wantErr:  false,
		},
		{
			name:     "SQL Server",
			dialect:  "sqlserver",
			wantType: &SQLServerDialect{},
			wantErr:  false,
		},
		{
			name:     "SQL Server (alternative name)",
			dialect:  "mssql",
			wantType: &SQLServerDialect{},
			wantErr:  false,
		},
		{
			name:     "Oracle",
			dialect:  "oracle",
			wantType: &OracleDialect{},
			wantErr:  false,
		},
		{
			name:     "Generic",
			dialect:  "generic",
			wantType: &GenericDialect{},
			wantErr:  false,
		},
		{
			name:     "Case insensitive",
			dialect:  "POSTGRES",
			wantType: &PostgresDialect{},
			wantErr:  false,
		},
		{
			name:     "Unsupported dialect",
			dialect:  "unsupported",
			wantType: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dialect, err := GetDialect(tt.dialect)
			
			// Check error
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDialect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
			// If we expect an error, no need to check the dialect type
			if tt.wantErr {
				return
			}
			
			// Check dialect type
			switch tt.wantType.(type) {
			case *PostgresDialect:
				if _, ok := dialect.(*PostgresDialect); !ok {
					t.Errorf("GetDialect() returned wrong type for %s", tt.dialect)
				}
			case *MySQLDialect:
				if _, ok := dialect.(*MySQLDialect); !ok {
					t.Errorf("GetDialect() returned wrong type for %s", tt.dialect)
				}
			case *SQLiteDialect:
				if _, ok := dialect.(*SQLiteDialect); !ok {
					t.Errorf("GetDialect() returned wrong type for %s", tt.dialect)
				}
			case *SQLServerDialect:
				if _, ok := dialect.(*SQLServerDialect); !ok {
					t.Errorf("GetDialect() returned wrong type for %s", tt.dialect)
				}
			case *OracleDialect:
				if _, ok := dialect.(*OracleDialect); !ok {
					t.Errorf("GetDialect() returned wrong type for %s", tt.dialect)
				}
			case *GenericDialect:
				if _, ok := dialect.(*GenericDialect); !ok {
					t.Errorf("GetDialect() returned wrong type for %s", tt.dialect)
				}
			default:
				t.Errorf("Unknown expected type for %s", tt.dialect)
			}
			
			// Check dialect name
			if dialect.Name() == "" {
				t.Errorf("Dialect name should not be empty")
			}
		})
	}
}

func TestBaseDialect_FormatValue(t *testing.T) {
	d := &BaseDialect{}
	
	tests := []struct {
		name  string
		value interface{}
		want  string
	}{
		{
			name:  "Nil value",
			value: nil,
			want:  "NULL",
		},
		{
			name:  "String value",
			value: "test",
			want:  "'test'",
		},
		{
			name:  "String with single quote",
			value: "test's",
			want:  "'test''s'",
		},
		{
			name:  "Integer value",
			value: 42,
			want:  "42",
		},
		{
			name:  "Float value",
			value: 3.14,
			want:  "3.14",
		},
		{
			name:  "Boolean true",
			value: true,
			want:  "TRUE",
		},
		{
			name:  "Boolean false",
			value: false,
			want:  "FALSE",
		},
		{
			name:  "Custom type",
			value: struct{ Name string }{"test"},
			want:  "'{test}'",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := d.FormatValue(tt.value)
			if got != tt.want {
				t.Errorf("BaseDialect.FormatValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColumnDef(t *testing.T) {
	// Test creating a ColumnDef
	col := ColumnDef{
		Name:     "id",
		Type:     SQLTypeInteger,
		Nullable: false,
	}
	
	if col.Name != "id" {
		t.Errorf("Expected column name 'id', got %s", col.Name)
	}
	
	if col.Type != SQLTypeInteger {
		t.Errorf("Expected column type INTEGER, got %s", col.Type)
	}
	
	if col.Nullable {
		t.Errorf("Expected column to be NOT NULL")
	}
	
	// Test SQLType constants
	if SQLTypeInteger != "INTEGER" {
		t.Errorf("Expected SQLTypeInteger to be 'INTEGER', got %s", SQLTypeInteger)
	}
	
	if SQLTypeFloat != "FLOAT" {
		t.Errorf("Expected SQLTypeFloat to be 'FLOAT', got %s", SQLTypeFloat)
	}
	
	if SQLTypeText != "TEXT" {
		t.Errorf("Expected SQLTypeText to be 'TEXT', got %s", SQLTypeText)
	}
	
	if SQLTypeDate != "DATE" {
		t.Errorf("Expected SQLTypeDate to be 'DATE', got %s", SQLTypeDate)
	}
	
	if SQLTypeDateTime != "DATETIME" {
		t.Errorf("Expected SQLTypeDateTime to be 'DATETIME', got %s", SQLTypeDateTime)
	}
	
	if SQLTypeBoolean != "BOOLEAN" {
		t.Errorf("Expected SQLTypeBoolean to be 'BOOLEAN', got %s", SQLTypeBoolean)
	}
}