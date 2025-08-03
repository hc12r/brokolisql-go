package errors

import (
	"fmt"
)

type ErrorType string

const (
	ErrorTypeInput     ErrorType = "INPUT_ERROR"
	ErrorTypeFormat    ErrorType = "FORMAT_ERROR"
	ErrorTypeTransform ErrorType = "TRANSFORM_ERROR"
	ErrorTypeSQL       ErrorType = "SQL_ERROR"
	ErrorTypeOutput    ErrorType = "OUTPUT_ERROR"
	ErrorTypeInternal  ErrorType = "INTERNAL_ERROR"
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
