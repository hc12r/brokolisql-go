package common

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type DataRow map[string]interface{}

type DataSet struct {
	Columns []string
	Rows    []DataRow
}

func ParseJSONData(jsonBytes []byte) ([]map[string]interface{}, error) {
	var data []map[string]interface{}
	err := json.Unmarshal(jsonBytes, &data)

	if err != nil || len(data) == 0 {

		var singleObject map[string]interface{}
		err = json.Unmarshal(jsonBytes, &singleObject)
		if err != nil {
			return nil, fmt.Errorf("failed to parse JSON data: %w", err)
		}
		data = []map[string]interface{}{singleObject}
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("no data found in JSON content")
	}

	return data, nil
}

func ConvertToDataSet(data []map[string]interface{}) *DataSet {

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

			if IsComplex(value) {
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
	}
}

func IsComplex(v interface{}) bool {
	if v == nil {
		return false
	}

	kind := reflect.TypeOf(v).Kind()
	return kind == reflect.Map || kind == reflect.Slice || kind == reflect.Array
}
