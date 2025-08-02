package errors

import (
	"fmt"
)

type ErrorType string

const (
	// ErrorTypeInput represents an input-related error
	ErrorTypeInput ErrorType = "INPUT_ERROR"

	// ErrorTypeFormat represents a format-related error
	ErrorTypeFormat ErrorType = "FORMAT_ERROR"

	// ErrorTypeTransform represents a transformation-related error
	ErrorTypeTransform ErrorType = "TRANSFORM_ERROR"

	// ErrorTypeSQL represents an SQL generation error
	ErrorTypeSQL ErrorType = "SQL_ERROR"

	// ErrorTypeOutput represents an output-related error
	ErrorTypeOutput ErrorType = "OUTPUT_ERROR"

	// ErrorTypeInternal represents an internal error
	ErrorTypeInternal ErrorType = "INTERNAL_ERROR"
)

type AppError struct {
	Type    ErrorType
	Message string
	Cause   error
}

func (e *AppError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s (cause: %v)", e.Type, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Cause
}

func NewInputError(message string, cause error) *AppError {
	return &AppError{
		Type:    ErrorTypeInput,
		Message: message,
		Cause:   cause,
	}
}

func NewFormatError(message string, cause error) *AppError {
	return &AppError{
		Type:    ErrorTypeFormat,
		Message: message,
		Cause:   cause,
	}
}

func NewTransformError(message string, cause error) *AppError {
	return &AppError{
		Type:    ErrorTypeTransform,
		Message: message,
		Cause:   cause,
	}
}

func NewSQLError(message string, cause error) *AppError {
	return &AppError{
		Type:    ErrorTypeSQL,
		Message: message,
		Cause:   cause,
	}
}

func NewOutputError(message string, cause error) *AppError {
	return &AppError{
		Type:    ErrorTypeOutput,
		Message: message,
		Cause:   cause,
	}
}

func NewInternalError(message string, cause error) *AppError {
	return &AppError{
		Type:    ErrorTypeInternal,
		Message: message,
		Cause:   cause,
	}
}
