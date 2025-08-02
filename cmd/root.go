package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"brokolisql-go/pkg/loaders"
	"brokolisql-go/pkg/services"
	"brokolisql-go/pkg/transformers"
	"brokolisql-go/pkg/utils"
)

var (
	inputFile        string
	outputFile       string
	tableName        string
	format           string
	dialect          string
	batchSize        int
	createTable      bool
	transformFile    string
	normalizeColumns bool
)

var rootCmd = &cobra.Command{
	Use:   "brokolisql",
	Short: "BrokoliSQL converts structured data files to SQL INSERT statements",
	Long: `BrokoliSQL is a command-line tool designed to facilitate the conversion of 
structured data files—such as CSV, Excel, JSON, and XML—into SQL INSERT statements.

It solves common problems faced during data import, transformation, and database 
seeding by offering a flexible, extensible, and easy-to-use interface.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runConversion()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	flags := rootCmd.PersistentFlags()
	flags.StringVar(&inputFile, "input", "", "Input file path (required)")
	flags.StringVar(&outputFile, "output", "", "Output SQL file path (required)")
	flags.StringVar(&tableName, "table", "", "Table name for SQL statements (required)")
	flags.StringVar(&format, "format", "", "Input file format (csv, json, xml, xlsx) - if not specified, will be inferred from file extension")
	flags.StringVar(&dialect, "dialect", "generic", "SQL dialect (generic, postgres, mysql, sqlite, sqlserver, oracle)")
	flags.IntVar(&batchSize, "batch-size", 100, "Number of rows per INSERT statement")
	flags.BoolVar(&createTable, "create-table", false, "Generate CREATE TABLE statement")
	flags.StringVar(&transformFile, "transform", "", "JSON file with transformation rules")
	flags.BoolVar(&normalizeColumns, "normalize", true, "Normalize column names for SQL compatibility")

	flags.StringVarP(&inputFile, "i", "i", "", "Input file path (shorthand)")
	flags.StringVarP(&outputFile, "o", "o", "", "Output SQL file path (shorthand)")
	flags.StringVarP(&tableName, "t", "t", "", "Table name for SQL statements (shorthand)")
	flags.StringVarP(&format, "f", "f", "", "Input file format (shorthand)")
	flags.StringVarP(&dialect, "d", "d", "generic", "SQL dialect (shorthand)")
	flags.IntVarP(&batchSize, "b", "b", 100, "Number of rows per INSERT statement (shorthand)")
	flags.BoolVarP(&createTable, "c", "c", false, "Generate CREATE TABLE statement (shorthand)")
	flags.StringVarP(&transformFile, "r", "r", "", "JSON file with transformation rules (shorthand)")
	flags.BoolVarP(&normalizeColumns, "n", "n", true, "Normalize column names for SQL compatibility (shorthand)")

	utils.CheckError(rootCmd.MarkFlagRequired("input"))
	utils.CheckError(rootCmd.MarkFlagRequired("output"))
	utils.CheckError(rootCmd.MarkFlagRequired("table"))

}

func runConversion() error {

	if format == "" {
		ext := filepath.Ext(inputFile)
		switch ext {
		case ".csv":
			format = "csv"
		case ".json":
			format = "json"
		case ".xml":
			format = "xml"
		case ".xlsx", ".xls":
			format = "excel"
		default:
			return fmt.Errorf("could not determine file format from extension: %s, please specify with --format", ext)
		}
	}

	loader, err := loaders.GetLoader(inputFile)
	if err != nil {
		return fmt.Errorf("failed to get loader: %w", err)
	}

	dataset, err := loader.Load(inputFile)
	if err != nil {
		return fmt.Errorf("failed to load data: %w", err)
	}

	if transformFile != "" {
		transformEngine, err := transformers.NewTransformEngine(transformFile)
		if err != nil {
			return fmt.Errorf("failed to initialize transform engine: %w", err)
		}

		if err := transformEngine.ApplyTransformations(dataset); err != nil {
			return fmt.Errorf("failed to apply transformations: %w", err)
		}
	}

	sqlGenerator, err := services.NewSQLGenerator(services.SQLGeneratorOptions{
		Dialect:          dialect,
		TableName:        tableName,
		CreateTable:      createTable,
		BatchSize:        batchSize,
		NormalizeColumns: normalizeColumns,
	})
	if err != nil {
		return fmt.Errorf("failed to initialize SQL generator: %w", err)
	}

	sql, err := sqlGenerator.Generate(dataset)
	if err != nil {
		return fmt.Errorf("failed to generate SQL: %w", err)
	}

	if err := os.WriteFile(outputFile, []byte(sql), 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	fmt.Printf("Successfully converted %s to SQL and saved to %s\n", inputFile, outputFile)
	return nil
}
