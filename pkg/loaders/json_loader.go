package loaders

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type JSONLoader struct{}

func (l *JSONLoader) Load(filePath string) (*DataSet, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read JSON file: %w", err)
	}

	var data []map[string]interface{}
	err = json.Unmarshal(file, &data)

	if err != nil || len(data) == 0 {

		var singleObject map[string]interface{}
		err = json.Unmarshal(file, &singleObject)
		if err != nil {
			return nil, fmt.Errorf("failed to parse JSON data: %w", err)
		}
		data = []map[string]interface{}{singleObject}
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("no data found in JSON file")
	}

	columnSet := make(map[string]bool)
	for _, obj := range data {
		for key := range obj {
			columnSet[key] = true
		}
	}

	columns := make([]string, 0, len(columnSet))
	for col := range columnSet {
		columns = append(columns, col)
	}

	rows := make([]DataRow, 0, len(data))
	for _, obj := range data {
		row := make(DataRow)
		for key, value := range obj {

			if isComplex(value) {
				jsonBytes, err := json.Marshal(value)
				if err == nil {
					row[key] = string(jsonBytes)
				} else {
					row[key] = fmt.Sprintf("%v", value)
				}
			} else {
				row[key] = value
			}
		}
		rows = append(rows, row)
	}

	return &DataSet{
		Columns: columns,
		Rows:    rows,
	}, nil
}

func isComplex(v interface{}) bool {
	if v == nil {
		return false
	}

	kind := reflect.TypeOf(v).Kind()
	return kind == reflect.Map || kind == reflect.Slice || kind == reflect.Array
}
