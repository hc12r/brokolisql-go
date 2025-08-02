package services

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"brokolisql-go/pkg/dialects"
	"brokolisql-go/pkg/loaders"
)

type TypeInferenceEngine struct {
	DateFormats []string

	TypeThreshold float64
}

func NewTypeInferenceEngine() *TypeInferenceEngine {
	return &TypeInferenceEngine{
		DateFormats: []string{
			"2006-01-02",
			"2006/01/02",
			"01/02/2006",
			"01-02-2006",
			"02/01/2006",
			"02-01-2006",
			time.RFC3339,
		},
		TypeThreshold: 0.8, // 80% of values must match to infer a type
	}
}

func (e *TypeInferenceEngine) InferColumnTypes(columns []string, rows []loaders.DataRow) map[string]dialects.SQLType {
	columnTypes := make(map[string]dialects.SQLType)

	for _, col := range columns {

		var values []interface{}
		for _, row := range rows {
			if val, ok := row[col]; ok && val != nil {
				values = append(values, val)
			}
		}

		columnTypes[col] = e.inferType(values)
	}

	return columnTypes
}

func (e *TypeInferenceEngine) inferType(values []interface{}) dialects.SQLType {
	if len(values) == 0 {
		return dialects.SQLTypeText // Default to TEXT for empty columns
	}

	intCount := 0
	floatCount := 0
	boolCount := 0
	dateCount := 0
	dateTimeCount := 0

	for _, val := range values {
		switch v := val.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			intCount++
		case float32, float64:
			floatCount++
		case bool:
			boolCount++
		case string:

			if e.isInteger(v) {
				intCount++
			} else if e.isFloat(v) {
				floatCount++
			} else if e.isBoolean(v) {
				boolCount++
			} else if _, isDate, hasTime := e.isDateTime(v); isDate {
				if hasTime {
					dateTimeCount++
				} else {
					dateCount++
				}
			}
		}
	}

	total := float64(len(values))
	intPercent := float64(intCount) / total
	floatPercent := float64(floatCount) / total
	boolPercent := float64(boolCount) / total
	datePercent := float64(dateCount) / total
	dateTimePercent := float64(dateTimeCount) / total

	if boolPercent >= e.TypeThreshold {
		return dialects.SQLTypeBoolean
	} else if intPercent >= e.TypeThreshold {
		return dialects.SQLTypeInteger
	} else if (intPercent + floatPercent) >= e.TypeThreshold {
		return dialects.SQLTypeFloat
	} else if dateTimePercent >= e.TypeThreshold {
		return dialects.SQLTypeDateTime
	} else if datePercent >= e.TypeThreshold {
		return dialects.SQLTypeDate
	}

	return dialects.SQLTypeText
}

func (e *TypeInferenceEngine) isInteger(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}

	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func (e *TypeInferenceEngine) isFloat(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}

	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func (e *TypeInferenceEngine) isBoolean(s string) bool {
	s = strings.ToLower(strings.TrimSpace(s))
	return s == "true" || s == "false" || s == "yes" || s == "no" || s == "1" || s == "0" || s == "t" || s == "f" || s == "y" || s == "n"
}

func (e *TypeInferenceEngine) isDateTime(s string) (time.Time, bool, bool) {
	s = strings.TrimSpace(s)
	if s == "" {
		return time.Time{}, false, false
	}

	hasTime := regexp.MustCompile(`\d{1,2}:\d{1,2}`).MatchString(s)

	for _, format := range e.DateFormats {
		if t, err := time.Parse(format, s); err == nil {
			return t, true, hasTime
		}
	}

	return time.Time{}, false, false
}
