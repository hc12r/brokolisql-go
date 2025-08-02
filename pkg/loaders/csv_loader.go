package loaders

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type CSVLoader struct{}

func (l *CSVLoader) Load(filePath string) (*DataSet, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV headers: %w", err)
	}

	for i, header := range headers {
		headers[i] = strings.TrimSpace(header)
	}

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV data: %w", err)
	}

	rows := make([]DataRow, 0, len(records))
	for _, record := range records {
		row := make(DataRow)
		for i, value := range record {
			if i < len(headers) {
				row[headers[i]] = value
			}
		}
		rows = append(rows, row)
	}

	return &DataSet{
		Columns: headers,
		Rows:    rows,
	}, nil
}
