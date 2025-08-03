package main

import (
	"brokolisql-go/internal/processing"
	"brokolisql-go/internal/transformers"
	"brokolisql-go/pkg/fetchers"
	"flag"
	"os"
	"path/filepath"

	"brokolisql-go/pkg/common"
	"brokolisql-go/pkg/loaders"
)

func main() {
	// Define command-line flags
	inputFile := flag.String("input", "", "Input file path (required unless using fetch mode)")
	outputFile := flag.String("output", "", "Output SQL file path (required)")
	tableName := flag.String("table", "", "Table name for SQL statements (required)")
	format := flag.String("format", "", "Input file format (csv, json, xml, xlsx) - if not specified, will be inferred from file extension")
	dialect := flag.String("dialect", "generic", "SQL dialect (generic, postgres, mysql, sqlite, sqlserver, oracle)")
	batchSize := flag.Int("batch-size", 100, "Number of rows per INSERT statement")
	createTable := flag.Bool("create-table", false, "Generate CREATE TABLE statement")
	transformFile := flag.String("transform", "", "JSON file with transformation rules")
	normalizeColumns := flag.Bool("normalize", true, "Normalize column names for SQL compatibility")
	logLevel := flag.String("log-level", "info", "Log level (debug, info, warning, error, fatal)")

	// Fetch mode flags
	fetchMode := flag.Bool("fetch", false, "Enable fetch mode to retrieve data from remote sources")
	fetchSource := flag.String("source", "", "Source URL or connection string for fetch mode")
	fetchType := flag.String("source-type", "rest", "Source type for fetch mode (rest, etc.)")

	// Parse flags
	flag.Parse()

	// Set up logger
	logger := common.NewLogger(common.LogLevelFromString(*logLevel))
	logger.Info("Starting BrokoliSQL")

	var dataset *common.DataSet
	var err error

	// Check if we're in fetch mode or file mode
	if *fetchMode {
		// Validate fetch mode parameters
		if *fetchSource == "" {
			logger.Fatal("Source URL or connection string is required when using fetch mode")
		}

		logger.Info("Fetch mode enabled, retrieving data from %s using %s fetcher", *fetchSource, *fetchType)

		// Get the appropriate fetcher
		fetcher, err := fetchers.GetFetcher(*fetchType)
		if err != nil {
			logger.Fatal("Failed to get fetcher: %v", err)
		}

		// Create options map for the fetcher
		options := make(map[string]interface{})
		// Add default options for REST fetcher
		if *fetchType == "rest" {
			options["method"] = "GET"
			options["headers"] = map[string]string{
				"Accept": "application/json",
			}
		}

		// Fetch the data
		dataset, err = fetcher.Fetch(*fetchSource, options)
		if err != nil {
			logger.Fatal("Failed to fetch data: %v", err)
		}

		logger.Info("Successfully fetched %d rows of data", len(dataset.Rows))
	} else {
		// Traditional file loading mode
		// Validate required flags
		if *inputFile == "" || *outputFile == "" || *tableName == "" {
			logger.Fatal("Input, output, and table flags are required when not using fetch mode")
		}

		// Determine file format if not specified
		fileFormat := *format
		if fileFormat == "" {
			ext := filepath.Ext(*inputFile)
			switch ext {
			case ".csv":
				fileFormat = "csv"
			case ".json":
				fileFormat = "json"
			case ".xml":
				fileFormat = "xml"
			case ".xlsx", ".xls":
				fileFormat = "excel"
			default:
				logger.Fatal("Could not determine file format from extension: %s, please specify with --format", ext)
			}
		}

		logger.Info("Processing file: %s (format: %s)", *inputFile, fileFormat)

		// Get the appropriate loader
		loader, err := loaders.GetLoader(*inputFile)
		if err != nil {
			logger.Fatal("Failed to get loader: %v", err)
		}

		// Load the data
		logger.Info("Loading data from file")
		dataset, err = loader.Load(*inputFile)
		if err != nil {
			logger.Fatal("Failed to load data: %v", err)
		}
	}

	logger.Info("Loaded %d rows with %d columns", len(dataset.Rows), len(dataset.Columns))

	// Apply transformations if specified
	if *transformFile != "" {
		logger.Info("Applying transformations from %s", *transformFile)
		transformEngine, err := transformers.NewTransformEngine(*transformFile)
		if err != nil {
			logger.Fatal("Failed to initialize transform engine: %v", err)
		}

		if err := transformEngine.ApplyTransformations(dataset); err != nil {
			logger.Fatal("Failed to apply transformations: %v", err)
		}

		logger.Info("Transformations applied successfully, resulting in %d rows", len(dataset.Rows))
	}

	// Generate SQL
	logger.Info("Generating SQL with dialect: %s", *dialect)
	sqlGenerator, err := processing.NewSQLGenerator(processing.SQLGeneratorOptions{
		Dialect:          *dialect,
		TableName:        *tableName,
		CreateTable:      *createTable,
		BatchSize:        *batchSize,
		NormalizeColumns: *normalizeColumns,
	})
	if err != nil {
		logger.Fatal("Failed to initialize SQL generator: %v", err)
	}

	sql, err := sqlGenerator.Generate(dataset)
	if err != nil {
		logger.Fatal("Failed to generate SQL: %v", err)
	}

	// Write output
	logger.Info("Writing SQL to %s", *outputFile)
	if err := os.WriteFile(*outputFile, []byte(sql), 0644); err != nil {
		logger.Fatal("Failed to write output file: %v", err)
	}

	logger.Info("Successfully converted %s to SQL and saved to %s", *inputFile, *outputFile)
}
