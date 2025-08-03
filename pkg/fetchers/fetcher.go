package fetchers

import (
	"brokolisql-go/pkg/common"
	"errors"
)

var (
	ErrUnsupportedSourceType = errors.New("unsupported source type")
)

type Fetcher interface {
	Fetch(source string, options map[string]interface{}) (*common.DataSet, error)
}

func GetFetcher(sourceType string) (Fetcher, error) {
	switch sourceType {
	case "rest":
		return &RESTFetcher{}, nil
	// Future implementations can be added here
	// case "database":
	//     return &DatabaseFetcher{}, nil
	default:
		return nil, ErrUnsupportedSourceType
	}
}
