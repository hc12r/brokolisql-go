# Nested JSON Support

BrokoliSQL-Go now supports processing JSON data with nested objects and arrays, automatically creating normalized SQL tables with proper relationships.

## Overview

When processing JSON data with nested structures, BrokoliSQL-Go will:

1. Detect nested objects and arrays within the JSON
2. Create separate tables for each nested structure
3. Establish foreign key relationships between parent and child tables
4. Maintain the correct order for table creation and data insertion
5. Generate SQL that respects these relationships

## Example

Given this JSON structure:

```json
{
  "id": 1,
  "name": "Alice",
  "address": {
    "city": "Maputo",
    "geo": {
      "lat": "-25.9",
      "lng": "32.6"
    }
  }
}
```

BrokoliSQL-Go will generate SQL that:

1. Creates a `geo` table with `id`, `lat`, and `lng` columns
2. Creates an `address` table with `id`, `city`, and `geo_id` columns (with a foreign key to `geo.id`)
3. Creates a `users` table with `id`, `name`, and `address_id` columns (with a foreign key to `address.id`)
4. Inserts data in the correct order: `geo` → `address` → `users`

## Usage

### Command Line

Use BrokoliSQL-Go as normal with JSON input:

```bash
brokolisql --input data.json --output output.sql --table users --create-table
```

The tool will automatically detect nested structures and handle them appropriately.

### Programmatic Usage

To use the nested JSON processing in your Go code:

1. Create a processor with default options:

```
// Create options
options := SQLGeneratorOptions{
    Dialect:          "generic",
    TableName:        "users",
    CreateTable:      true,
    BatchSize:        100,
    NormalizeColumns: true,
}

// Create the processor
processor, err := NewNestedJSONProcessor(options)
if err != nil {
    // Handle error
}

// Process the JSON data
sql, err := processor.ProcessNestedJSON(jsonData)
if err != nil {
    // Handle error
}

// Use the generated SQL
fmt.Println(sql)
```

## Customization

You can customize how nested JSON is processed using custom options:

```
// Create custom options
customOptions := NestedJSONProcessorOptions{
    SQLGeneratorOptions: SQLGeneratorOptions{
        Dialect:          "generic",
        TableName:        "users",
        CreateTable:      true,
        BatchSize:        100,
        NormalizeColumns: true,
    },
    NamingConvention: CamelCase,  // Use camelCase for table and column names
    TablePrefix:      "app_",     // Prefix all table names with "app_"
    PluralizeTable:   true,       // Use plural table names (e.g., "users" instead of "user")
}

// Create the processor with custom options
processor, err := NewNestedJSONProcessorWithOptions(customOptions)
```

### Naming Conventions

The following naming conventions are supported:

- `SnakeCase`: Uses snake_case for all names (default)
- `CamelCase`: Uses camelCase for all names
- `PascalCase`: Uses PascalCase for all names

### Array Handling

- Arrays of primitive values (strings, numbers, booleans) are stored as JSON strings in the parent table
- Arrays of objects are normalized into separate child tables with foreign keys back to the parent

## Limitations

- Currently, circular references in JSON are not supported
- Very deep nesting levels (>100) may cause performance issues
- The automatic schema inference is based on the structure of the provided data and may not capture all possible variations in a large dataset