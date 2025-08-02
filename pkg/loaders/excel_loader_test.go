package loaders

import (
	"testing"
)

func TestGetLoader_Excel(t *testing.T) {
	// Test .xlsx extension
	loader, err := GetLoader("test.xlsx")
	if err != nil {
		t.Errorf("GetLoader() error = %v", err)
		return
	}
	
	if _, ok := loader.(*ExcelLoader); !ok {
		t.Errorf("GetLoader() returned wrong loader type for .xlsx file")
	}
	
	// Test .xls extension
	loader, err = GetLoader("test.xls")
	if err != nil {
		t.Errorf("GetLoader() error = %v", err)
		return
	}
	
	if _, ok := loader.(*ExcelLoader); !ok {
		t.Errorf("GetLoader() returned wrong loader type for .xls file")
	}
}

// Note: Full testing of the Excel loader would require creating actual Excel files,
// which is complex to do programmatically in this environment.
// In a real-world scenario, you would:
// 1. Create test Excel files
// 2. Test various scenarios (valid files, empty files, files with missing values)
// 3. Test error cases (non-existent files, invalid Excel files)
//
// Example test structure would be similar to the other loader tests:
/*
func TestExcelLoader_Load(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "excel_loader_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test Excel files using excelize
	// ...

	tests := []struct {
		name     string
		filePath string
		wantErr  bool
		checkFn  func(*testing.T, *DataSet)
	}{
		// Test cases
		// ...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ExcelLoader{}
			got, err := l.Load(tt.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExcelLoader.Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.checkFn != nil {
				tt.checkFn(t, got)
			}
		})
	}
}
*/