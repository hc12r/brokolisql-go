package services

import (
	"reflect"
	"testing"
)

func TestNewNormalizer(t *testing.T) {
	n := NewNormalizer()

	// Check default values
	if n.MaxLength != 64 {
		t.Errorf("NewNormalizer() MaxLength = %v, want %v", n.MaxLength, 64)
	}

	if n.PreserveCase != false {
		t.Errorf("NewNormalizer() PreserveCase = %v, want %v", n.PreserveCase, false)
	}

	if n.ReplaceSpaces != true {
		t.Errorf("NewNormalizer() ReplaceSpaces = %v, want %v", n.ReplaceSpaces, true)
	}

	if n.SpaceReplacement != "_" {
		t.Errorf("NewNormalizer() SpaceReplacement = %v, want %v", n.SpaceReplacement, "_")
	}
}

func TestNormalizer_NormalizeColumnName(t *testing.T) {
	tests := []struct {
		name       string
		normalizer *Normalizer
		columnName string
		want       string
	}{
		{
			name:       "Simple column name",
			normalizer: NewNormalizer(),
			columnName: "column",
			want:       "COLUMN",
		},
		{
			name:       "Column name with spaces",
			normalizer: NewNormalizer(),
			columnName: "column name",
			want:       "COLUMN_NAME",
		},
		{
			name:       "Column name with special characters",
			normalizer: NewNormalizer(),
			columnName: "column-name!",
			want:       "COLUMN_NAME_",
		},
		{
			name:       "Column name starting with number",
			normalizer: NewNormalizer(),
			columnName: "123column",
			want:       "_123COLUMN",
		},
		{
			name:       "Column name with leading/trailing spaces",
			normalizer: NewNormalizer(),
			columnName: "  column  ",
			want:       "COLUMN",
		},
		{
			name:       "Empty column name",
			normalizer: NewNormalizer(),
			columnName: "",
			want:       "COLUMN",
		},
		{
			name: "Column name with custom settings",
			normalizer: &Normalizer{
				MaxLength:        10,
				PreserveCase:     true,
				ReplaceSpaces:    true,
				SpaceReplacement: "-",
			},
			columnName: "This is a very long column name",
			want:       "This-is-a-",
		},
		{
			name: "Column name without space replacement",
			normalizer: &Normalizer{
				MaxLength:     64,
				PreserveCase:  false,
				ReplaceSpaces: false,
			},
			columnName: "column name",
			want:       "COLUMN NAME",
		},
		{
			name: "Column name with custom space replacement",
			normalizer: &Normalizer{
				MaxLength:        64,
				PreserveCase:     false,
				ReplaceSpaces:    true,
				SpaceReplacement: ".",
			},
			columnName: "column name",
			want:       "COLUMN.NAME",
		},
		{
			name: "Column name with case preservation",
			normalizer: &Normalizer{
				MaxLength:        64,
				PreserveCase:     true,
				ReplaceSpaces:    true,
				SpaceReplacement: "_",
			},
			columnName: "Column Name",
			want:       "Column_Name",
		},
		{
			name: "Column name with max length",
			normalizer: &Normalizer{
				MaxLength:        5,
				PreserveCase:     false,
				ReplaceSpaces:    true,
				SpaceReplacement: "_",
			},
			columnName: "column_name",
			want:       "COLUM",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.normalizer.NormalizeColumnName(tt.columnName)
			if got != tt.want {
				t.Errorf("Normalizer.NormalizeColumnName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizer_NormalizeColumnNames(t *testing.T) {
	tests := []struct {
		name        string
		normalizer  *Normalizer
		columnNames []string
		want        []string
	}{
		{
			name:        "Simple column names",
			normalizer:  NewNormalizer(),
			columnNames: []string{"id", "name", "email"},
			want:        []string{"ID", "NAME", "EMAIL"},
		},
		{
			name:        "Column names with spaces",
			normalizer:  NewNormalizer(),
			columnNames: []string{"user id", "full name", "email address"},
			want:        []string{"USER_ID", "FULL_NAME", "EMAIL_ADDRESS"},
		},
		{
			name:        "Duplicate column names",
			normalizer:  NewNormalizer(),
			columnNames: []string{"id", "id", "id"},
			want:        []string{"ID", "ID_0", "ID_1"},
		},
		{
			name:        "Mixed column names",
			normalizer:  NewNormalizer(),
			columnNames: []string{"id", "user-name", "123email", ""},
			want:        []string{"ID", "USER_NAME", "_123EMAIL", "COLUMN"},
		},
		{
			name: "Column names with custom settings",
			normalizer: &Normalizer{
				MaxLength:        5,
				PreserveCase:     true,
				ReplaceSpaces:    true,
				SpaceReplacement: "-",
			},
			columnNames: []string{"user id", "full name", "email"},
			want:        []string{"user-", "full-", "email"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.normalizer.NormalizeColumnNames(tt.columnNames)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Normalizer.NormalizeColumnNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizer_NormalizeColumnNames_DuplicateHandling(t *testing.T) {
	n := NewNormalizer()

	// Test case where normalization creates duplicates
	columnNames := []string{"user-id", "user_id", "userId"}
	normalized := n.NormalizeColumnNames(columnNames)

	// Check that all normalized names are unique
	uniqueNames := make(map[string]bool)
	for _, name := range normalized {
		if uniqueNames[name] {
			t.Errorf("Duplicate normalized name found: %s", name)
		}
		uniqueNames[name] = true
	}

	// Check that we have the same number of names
	if len(normalized) != len(columnNames) {
		t.Errorf("NormalizeColumnNames() returned %d names, want %d", len(normalized), len(columnNames))
	}
}
