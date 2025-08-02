package loaders

import (
	"errors"
	"path/filepath"
)

type DataRow map[string]interface{}

type DataSet struct {
	Columns []string
	Rows    []DataRow
}

type Loader interface {
	Load(filePath string) (*DataSet, error)
}

func GetLoader(filePath string) (Loader, error) {
	ext := filepath.Ext(filePath)

	switch ext {
	case ".csv":
		return &CSVLoader{}, nil
	case ".json":
		return &JSONLoader{}, nil
	case ".xml":
		return &XMLLoader{}, nil
	case ".xlsx", ".xls":
		return &ExcelLoader{}, nil
	default:
		return nil, errors.New("unsupported file format: " + ext)
	}
}
