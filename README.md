# BrokoliSQL-Go

BrokoliSQL-Go is a powerful command-line tool written in Go that converts structured data files (CSV, Excel, JSON, XML) into SQL INSERT statements. It provides flexible data transformation capabilities and supports multiple SQL dialects, making it ideal for database seeding, data migration, and ETL workflows.

![BrokoliSQL-Go](https://img.shields.io/badge/BrokoliSQL-Go-brightgreen)

## Features

- **Multi-format Support**: Process CSV, Excel (XLSX), JSON, and XML files
- **SQL Dialect Support**: Generate SQL for PostgreSQL, MySQL, SQLite, SQL Server, Oracle, and more
- **Automatic Table Creation**: Optionally generate CREATE TABLE statements based on input data
- **Smart Type Inference**: Automatically detect appropriate SQL data types
- **Batch Processing**: Control the number of rows per INSERT statement for optimal performance
- **Powerful Transformations**: Apply various transformations to your data before SQL generation
- **Column Normalization**: Automatically normalize column names for SQL compatibility

## Installation

### Prerequisites

- Go 1.18 or higher

### From Source

```bash
# Clone the repository
git clone https://github.com/yourusername/brokolisql-go.git
cd brokolisql-go

# Build the binary
go build -o brokolisql

# Install globally (optional)
go install
```

## Usage

### Basic Usage

```bash
brokolisql --input data.csv --output output.sql --table users
```

### Command-Line Options

```
Usage:
  brokolisql [flags]

Flags:
  -b, --batch-size int       Number of rows per INSERT statement (default 100)
  -c, --create-table         Generate CREATE TABLE statement
  -d, --dialect string       SQL dialect (generic, postgres, mysql, sqlite, sqlserver, oracle) (default "generic")
  -f, --format string        Input file format (csv, json, xml, xlsx) - if not specified, will be inferred from file extension
  -h, --help                 help for brokolisql
  -i, --input string         Input file path (required)
  -n, --normalize            Normalize column names for SQL compatibility (default true)
  -o, --output string        Output SQL file path (required)
  -r, --transform string     JSON file with transformation rules
  -t, --table string         Table name for SQL statements (required)
      --log-level string     Log level (debug, info, warning, error, fatal) (default "info")
```

### Examples

Generate SQL with a specific dialect:

```bash
brokolisql --input data.csv --output output.sql --table users --dialect mysql
```

Generate a CREATE TABLE statement:

```bash
brokolisql --input data.csv --output output.sql --table users --create-table
```

Use batch inserts for better performance:

```bash
brokolisql --input data.csv --output output.sql --table users --batch-size 100
```

Apply transformations:

```bash
brokolisql --input data.csv --output output.sql --table users --transform transforms.json
```

## Data Transformations

BrokoliSQL-Go supports powerful data transformations through a JSON configuration file. Here's an example:

```json
{
  "transformations": [
    {
      "type": "rename_columns",
      "mapping": {
        "FIRST_NAME": "GIVEN_NAME",
        "LAST_NAME": "SURNAME"
      }
    },
    {
      "type": "add_column",
      "name": "FULL_NAME",
      "expression": "GIVEN_NAME + ' ' + SURNAME"
    },
    {
      "type": "filter_rows",
      "condition": "COUNTRY in ['USA', 'Canada', 'UK', 'Germany']"
    },
    {
      "type": "apply_function",
      "column": "EMAIL",
      "function": "lower"
    },
    {
      "type": "replace_values",
      "column": "COUNTRY",
      "mapping": {
        "USA": "United States",
        "UK": "United Kingdom"
      }
    },
    {
      "type": "drop_columns",
      "columns": ["TEMP_COLUMN"]
    },
    {
      "type": "sort",
      "columns": ["COUNTRY", "CITY"],
      "ascending": true
    }
  ]
}
```

## Use Cases

BrokoliSQL-Go is particularly useful in the following scenarios:

### Data Migration

When migrating data between systems, BrokoliSQL-Go can transform source data into SQL statements compatible with the target database, handling format conversions and data transformations in a single step.

### Database Seeding

For development and testing environments, BrokoliSQL-Go makes it easy to convert sample data from various formats into SQL for database initialization.

### ETL Workflows

As part of Extract-Transform-Load (ETL) pipelines, BrokoliSQL-Go can transform data from various sources and prepare it for loading into a database.

### Data Analysis

Data analysts can use BrokoliSQL-Go to quickly convert data from various formats into SQL for further analysis in a database environment.

### API Integration

When integrating with APIs that provide data in JSON or XML format, BrokoliSQL-Go can transform this data into SQL for storage in a relational database.

## Project Structure

```
brokolisql-go/
├── cmd/
│   └── root.go
├── examples/
│   ├── customers.csv
│   ├── output.sql
│   └── transforms.json
├── pkg/
│   ├── dialects/
│   │   ├── dialect.go
│   │   ├── generic.go
│   │   ├── mysql.go
│   │   ├── oracle.go
│   │   ├── postgres.go
│   │   ├── sqlite.go
│   │   └── sqlserver.go
│   ├── loaders/
│   │   ├── csv_loader.go
│   │   ├── excel_loader.go
│   │   ├── json_loader.go
│   │   ├── loader.go
│   │   └── xml_loader.go
│   ├── services/
│   │   ├── normalizer.go
│   │   ├── sql_generator.go
│   │   └── type_inference.go
│   ├── transformers/
│   │   └── transform_engine.go
│   └── utils/
│       ├── errors.go
│       └── logger.go
├── go.mod
├── go.sum
└── main.go
```

## Contributing

Contributions to BrokoliSQL-Go are welcome! Here's how you can contribute:

1. **Fork the Repository**: Create your own fork of the project.

2. **Create a Feature Branch**:
   ```bash
   git checkout -b feature/your-feature-name
   ```

3. **Make Your Changes**: Implement your feature or bug fix.

4. **Write Tests**: Add tests for your changes to ensure they work correctly.

5. **Run Tests**:
   ```bash
   go test ./...
   ```

6. **Format Your Code**:
   ```bash
   go fmt ./...
   ```

7. **Commit Your Changes**:
   ```bash
   git commit -m "Add feature: your feature description"
   ```

8. **Push to Your Fork**:
   ```bash
   git push origin feature/your-feature-name
   ```

9. **Create a Pull Request**: Open a pull request from your fork to the main repository.

### Development Guidelines

- Follow Go best practices and coding conventions
- Write clear, concise commit messages
- Document your code with comments
- Add unit tests for new functionality
- Update documentation when necessary

## License

BrokoliSQL-Go is licensed under the GNU GPL-3.0 License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

BrokoliSQL-Go is a Go implementation of the original [BrokoliSQL](https://github.com/hc12r/brokolisql) Python project, reimagined with Go's performance and concurrency benefits.