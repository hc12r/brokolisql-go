package dialects

import (
	"strings"
	"testing"
)

func TestGenericDialect_Name(t *testing.T) {
	d := &GenericDialect{}
	if got := d.Name(); got != "generic" {
		t.Errorf("GenericDialect.Name() = %v, want %v", got, "generic")
	}
}

func TestGenericDialect_QuoteIdentifier(t *testing.T) {
	d := &GenericDialect{}
	tests := []struct {
		name       string
		identifier string
		want       string
	}{
		{
			name:       "Simple identifier",
			identifier: "column",
			want:       "\"column\"",
		},
		{
			name:       "Identifier with spaces",
			identifier: "column name",
			want:       "\"column name\"",
		},
		{
			name:       "Identifier with special characters",
			identifier: "column-name",
			want:       "\"column-name\"",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := d.QuoteIdentifier(tt.identifier); got != tt.want {
				t.Errorf("GenericDialect.QuoteIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenericDialect_CreateTable(t *testing.T) {
	d := &GenericDialect{}
	
	tests := []struct {
		name      string
		tableName string
		columns   []ColumnDef
		contains  []string
	}{
		{
			name:      "Simple table",
			tableName: "users",
			columns: []ColumnDef{
				{Name: "id", Type: SQLTypeInteger, Nullable: false},
				{Name: "name", Type: SQLTypeText, Nullable: true},
			},
			contains: []string{
				"CREATE TABLE \"users\"",
				"\"id\" INTEGER NOT NULL",
				"\"name\" TEXT",
			},
		},
		{
			name:      "Table with various types",
			tableName: "products",
			columns: []ColumnDef{
				{Name: "id", Type: SQLTypeInteger, Nullable: false},
				{Name: "name", Type: SQLTypeText, Nullable: false},
				{Name: "price", Type: SQLTypeFloat, Nullable: true},
				{Name: "created_at", Type: SQLTypeDateTime, Nullable: true},
				{Name: "is_active", Type: SQLTypeBoolean, Nullable: false},
			},
			contains: []string{
				"CREATE TABLE \"products\"",
				"\"id\" INTEGER NOT NULL",
				"\"name\" TEXT NOT NULL",
				"\"price\" FLOAT",
				"\"created_at\" DATETIME",
				"\"is_active\" BOOLEAN NOT NULL",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql := d.CreateTable(tt.tableName, tt.columns)
			
			// Check that the SQL contains all expected strings
			for _, s := range tt.contains {
				if !strings.Contains(sql, s) {
					t.Errorf("GenericDialect.CreateTable() = %v, should contain %v", sql, s)
				}
			}
			
			// Check that the SQL ends with a semicolon and newline
			if !strings.HasSuffix(sql, ";\n") {
				t.Errorf("GenericDialect.CreateTable() = %v, should end with semicolon and newline", sql)
			}
		})
	}
}

func TestGenericDialect_InsertInto(t *testing.T) {
	d := &GenericDialect{}
	
	tests := []struct {
		name      string
		tableName string
		columns   []string
		values    [][]interface{}
		batchSize int
		contains  []string
	}{
		{
			name:      "Simple insert",
			tableName: "users",
			columns:   []string{"id", "name"},
			values: [][]interface{}{
				{1, "John"},
				{2, "Jane"},
			},
			batchSize: 0, // Use default batch size
			contains: []string{
				"INSERT INTO \"users\" (\"id\", \"name\") VALUES",
				"(1, 'John')",
				"(2, 'Jane')",
			},
		},
		{
			name:      "Insert with various types",
			tableName: "products",
			columns:   []string{"id", "name", "price", "is_active"},
			values: [][]interface{}{
				{1, "Product 1", 10.5, true},
				{2, "Product 2", 20.75, false},
			},
			batchSize: 0, // Use default batch size
			contains: []string{
				"INSERT INTO \"products\" (\"id\", \"name\", \"price\", \"is_active\") VALUES",
				"(1, 'Product 1', 10.5, TRUE)",
				"(2, 'Product 2', 20.75, FALSE)",
			},
		},
		{
			name:      "Insert with batch size",
			tableName: "users",
			columns:   []string{"id", "name"},
			values: [][]interface{}{
				{1, "John"},
				{2, "Jane"},
				{3, "Bob"},
				{4, "Alice"},
			},
			batchSize: 2, // Split into batches of 2
			contains: []string{
				"INSERT INTO \"users\" (\"id\", \"name\") VALUES",
				"(1, 'John')",
				"(2, 'Jane')",
				"INSERT INTO \"users\" (\"id\", \"name\") VALUES",
				"(3, 'Bob')",
				"(4, 'Alice')",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql := d.InsertInto(tt.tableName, tt.columns, tt.values, tt.batchSize)
			
			// Check that the SQL contains all expected strings
			for _, s := range tt.contains {
				if !strings.Contains(sql, s) {
					t.Errorf("GenericDialect.InsertInto() = %v, should contain %v", sql, s)
				}
			}
			
			// Check that the SQL ends with a semicolon and newlines
			if !strings.HasSuffix(sql, ";\n\n") {
				t.Errorf("GenericDialect.InsertInto() = %v, should end with semicolon and newlines", sql)
			}
			
			// Check batch size behavior
			if tt.batchSize > 0 {
				// Count the number of INSERT statements
				insertCount := strings.Count(sql, "INSERT INTO")
				expectedInsertCount := (len(tt.values) + tt.batchSize - 1) / tt.batchSize // Ceiling division
				if insertCount != expectedInsertCount {
					t.Errorf("GenericDialect.InsertInto() has %d INSERT statements, want %d", insertCount, expectedInsertCount)
				}
			}
		})
	}
}