package processing

import (
	"brokolisql-go/internal/dialects"
	"brokolisql-go/pkg/loaders"
)

type SQLGeneratorOptions struct {
	Dialect          string
	TableName        string
	CreateTable      bool
	BatchSize        int
	NormalizeColumns bool
}

type SQLGenerator struct {
	options     SQLGeneratorOptions
	normalizer  *Normalizer
	typeInferer *TypeInferenceEngine
	dialect     dialects.Dialect
}

func NewSQLGenerator(options SQLGeneratorOptions) (*SQLGenerator, error) {

	if options.TableName == "" {
		options.TableName = "data"
	}

	if options.BatchSize <= 0 {
		options.BatchSize = 100
	}

	dialect, err := dialects.GetDialect(options.Dialect)
	if err != nil {
		return nil, err
	}

	return &SQLGenerator{
		options:     options,
		normalizer:  NewNormalizer(),
		typeInferer: NewTypeInferenceEngine(),
		dialect:     dialect,
	}, nil
}

func (g *SQLGenerator) Generate(dataset *loaders.DataSet) (string, error) {

	columns := dataset.Columns
	if g.options.NormalizeColumns {
		columns = g.normalizer.NormalizeColumnNames(columns)
	}

	if g.options.NormalizeColumns {
		normalizedRows := make([]loaders.DataRow, len(dataset.Rows))
		for i, row := range dataset.Rows {
			normalizedRow := make(loaders.DataRow)
			for j, col := range dataset.Columns {
				normalizedCol := columns[j]
				normalizedRow[normalizedCol] = row[col]
			}
			normalizedRows[i] = normalizedRow
		}
		dataset.Rows = normalizedRows
	}

	var sql string

	if g.options.CreateTable {

		columnTypes := g.typeInferer.InferColumnTypes(columns, dataset.Rows)

		columnDefs := make([]dialects.ColumnDef, len(columns))
		for i, col := range columns {
			columnDefs[i] = dialects.ColumnDef{
				Name:     col,
				Type:     columnTypes[col],
				Nullable: true, // Default to nullable
			}
		}

		sql += g.dialect.CreateTable(g.options.TableName, columnDefs)
		sql += "\n"
	}

	values := make([][]interface{}, len(dataset.Rows))
	for i, row := range dataset.Rows {
		rowValues := make([]interface{}, len(columns))
		for j, col := range columns {
			rowValues[j] = row[col]
		}
		values[i] = rowValues
	}

	sql += g.dialect.InsertInto(g.options.TableName, columns, values, g.options.BatchSize)

	return sql, nil
}
