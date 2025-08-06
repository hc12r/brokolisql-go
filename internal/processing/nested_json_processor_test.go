package processing

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestNestedJSONProcessor_ProcessNestedJSON(t *testing.T) {
	// Test case from the instructions
	jsonData := `{
		"id": 1,
		"name": "Alice",
		"address": {
			"city": "Maputo",
			"geo": {
				"lat": "-25.9",
				"lng": "32.6"
			}
		}
	}`

	// Parse the JSON data
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	// Create a processor with default options
	options := SQLGeneratorOptions{
		Dialect:          "generic",
		TableName:        "users",
		CreateTable:      true,
		BatchSize:        100,
		NormalizeColumns: true,
	}
	processor, err := NewNestedJSONProcessor(options)
	if err != nil {
		t.Fatalf("Failed to create processor: %v", err)
	}

	// Process the data
	sql, err := processor.ProcessNestedJSON([]map[string]interface{}{data})
	if err != nil {
		t.Fatalf("Failed to process nested JSON: %v", err)
	}

	// Verify the SQL output
	verifySQL(t, sql, []string{
		"CREATE TABLE", "geos", "lat", "lng",
		"CREATE TABLE", "addresses", "city", "geo_id",
		"CREATE TABLE", "userses", "id", "name", "address_id",
		"FOREIGN KEY", "REFERENCES",
	})

	// Verify table creation order
	// Print the SQL for debugging
	t.Logf("Generated SQL:\n%s", sql)
	verifyTableOrder(t, sql, []string{"geos", "addresses", "userses"})
}

func TestNestedJSONProcessor_ProcessArrays(t *testing.T) {
	// Test case with arrays
	jsonData := `{
		"id": 1,
		"name": "Alice",
		"tags": ["developer", "golang"],
		"contacts": [
			{"type": "email", "value": "alice@example.com"},
			{"type": "phone", "value": "123-456-7890"}
		]
	}`

	// Parse the JSON data
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	// Create a processor with default options
	options := SQLGeneratorOptions{
		Dialect:          "generic",
		TableName:        "users",
		CreateTable:      true,
		BatchSize:        100,
		NormalizeColumns: true,
	}
	processor, err := NewNestedJSONProcessor(options)
	if err != nil {
		t.Fatalf("Failed to create processor: %v", err)
	}

	// Process the data
	sql, err := processor.ProcessNestedJSON([]map[string]interface{}{data})
	if err != nil {
		t.Fatalf("Failed to process nested JSON: %v", err)
	}

	// Verify the SQL output
	// Print the SQL for debugging
	t.Logf("Generated SQL:\n%s", sql)
	verifySQL(t, sql, []string{
		"CREATE TABLE", "contactses", "type", "value", "userses_id",
		"CREATE TABLE", "userses", "id", "name", "tags",
		"FOREIGN KEY", "REFERENCES",
	})

	// Verify that primitive arrays are stored as JSON strings
	if !strings.Contains(sql, "tags") {
		t.Errorf("SQL should contain 'tags' column for primitive array")
	}

	// Verify that object arrays are stored in separate tables
	if !strings.Contains(sql, "contactses") {
		t.Errorf("SQL should contain 'contactses' table for array of objects")
	}
}

func TestNestedJSONProcessor_DeepNesting(t *testing.T) {
	// Test case with deep nesting
	jsonData := `{
		"id": 1,
		"name": "Alice",
		"company": {
			"name": "Acme Inc",
			"department": {
				"name": "Engineering",
				"location": {
					"building": "HQ",
					"floor": 3
				}
			}
		}
	}`

	// Parse the JSON data
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	// Create a processor with default options
	options := SQLGeneratorOptions{
		Dialect:          "generic",
		TableName:        "users",
		CreateTable:      true,
		BatchSize:        100,
		NormalizeColumns: true,
	}
	processor, err := NewNestedJSONProcessor(options)
	if err != nil {
		t.Fatalf("Failed to create processor: %v", err)
	}

	// Process the data
	sql, err := processor.ProcessNestedJSON([]map[string]interface{}{data})
	if err != nil {
		t.Fatalf("Failed to process nested JSON: %v", err)
	}

	// Verify the SQL output
	// Print the SQL for debugging
	t.Logf("Generated SQL:\n%s", sql)
	verifySQL(t, sql, []string{
		"CREATE TABLE", "locations", "building", "floor",
		"CREATE TABLE", "departments", "name", "location_id",
		"CREATE TABLE", "companies", "name", "department_id",
		"CREATE TABLE", "userses", "id", "name", "company_id",
		"FOREIGN KEY", "REFERENCES",
	})

	// Verify table creation order (deepest first)
	verifyTableOrder(t, sql, []string{"locations", "departments", "companies", "userses"})
}

func TestNestedJSONProcessor_CustomNamingConvention(t *testing.T) {
	// Test case with custom naming convention
	jsonData := `{
		"id": 1,
		"firstName": "Alice",
		"homeAddress": {
			"cityName": "Maputo"
		}
	}`

	// Parse the JSON data
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	// Create a processor with custom options
	options := NestedJSONProcessorOptions{
		SQLGeneratorOptions: SQLGeneratorOptions{
			Dialect:          "generic",
			TableName:        "users",
			CreateTable:      true,
			BatchSize:        100,
			NormalizeColumns: true,
		},
		NamingConvention: CamelCase,
		TablePrefix:      "app_",
		PluralizeTable:   true,
	}
	processor, err := NewNestedJSONProcessorWithOptions(options)
	if err != nil {
		t.Fatalf("Failed to create processor: %v", err)
	}

	// Process the data
	sql, err := processor.ProcessNestedJSON([]map[string]interface{}{data})
	if err != nil {
		t.Fatalf("Failed to process nested JSON: %v", err)
	}

	// Print the SQL for debugging
	t.Logf("Generated SQL:\n%s", sql)

	// Verify the SQL output uses camelCase and has the prefix
	if !strings.Contains(sql, "app_homeAddresses") {
		t.Errorf("SQL should contain 'app_homeAddresses' table with camelCase and prefix")
	}
	if !strings.Contains(sql, "app_userses") {
		t.Errorf("SQL should contain 'app_userses' table with prefix")
	}
}

// Helper function to verify that SQL contains expected strings
func verifySQL(t *testing.T, sql string, expected []string) {
	for _, exp := range expected {
		if !strings.Contains(sql, exp) {
			t.Errorf("SQL should contain '%s'", exp)
		}
	}
}

// Helper function to verify table creation order
func verifyTableOrder(t *testing.T, sql string, expectedOrder []string) {
	lastPos := -1
	for _, tableName := range expectedOrder {
		createTablePos := strings.Index(sql, "CREATE TABLE \""+tableName+"\"")
		if createTablePos == -1 {
			t.Errorf("SQL should contain CREATE TABLE for '%s'", tableName)
			continue
		}

		if createTablePos < lastPos {
			t.Errorf("Table '%s' should be created before '%s'",
				expectedOrder[len(expectedOrder)-2], tableName)
		}

		lastPos = createTablePos
	}
}
