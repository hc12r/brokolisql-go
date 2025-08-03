package errors

import (
	"errors"
	"strings"
	"testing"
)

func TestAppError_Error(t *testing.T) {
	// Test with cause
	cause := errors.New("underlying error")
	err := &AppError{
		Type:    ErrorTypeInput,
		Message: "test message",
		Cause:   cause,
	}

	errStr := err.Error()
	if !strings.Contains(errStr, string(ErrorTypeInput)) {
		t.Errorf("Error() should contain error type, got: %s", errStr)
	}
	if !strings.Contains(errStr, "test message") {
		t.Errorf("Error() should contain message, got: %s", errStr)
	}
	if !strings.Contains(errStr, "underlying error") {
		t.Errorf("Error() should contain cause, got: %s", errStr)
	}

	// Test without cause
	err = &AppError{
		Type:    ErrorTypeInput,
		Message: "test message",
		Cause:   nil,
	}

	errStr = err.Error()
	if !strings.Contains(errStr, string(ErrorTypeInput)) {
		t.Errorf("Error() should contain error type, got: %s", errStr)
	}
	if !strings.Contains(errStr, "test message") {
		t.Errorf("Error() should contain message, got: %s", errStr)
	}
	if strings.Contains(errStr, "cause") {
		t.Errorf("Error() should not contain cause when nil, got: %s", errStr)
	}
}

func TestAppError_Unwrap(t *testing.T) {
	cause := errors.New("underlying error")
	err := &AppError{
		Type:    ErrorTypeInput,
		Message: "test message",
		Cause:   cause,
	}

	unwrapped := err.Unwrap()
	if unwrapped != cause {
		t.Errorf("Unwrap() = %v, want %v", unwrapped, cause)
	}

	// Test with nil cause
	err = &AppError{
		Type:    ErrorTypeInput,
		Message: "test message",
		Cause:   nil,
	}

	unwrapped = err.Unwrap()
	if unwrapped != nil {
		t.Errorf("Unwrap() = %v, want nil", unwrapped)
	}
}

func TestNewInputError(t *testing.T) {
	cause := errors.New("underlying error")
	err := NewInputError("test message", cause)

	if err.Type != ErrorTypeInput {
		t.Errorf("NewInputError() Type = %v, want %v", err.Type, ErrorTypeInput)
	}
	if err.Message != "test message" {
		t.Errorf("NewInputError() Message = %v, want %v", err.Message, "test message")
	}
	if err.Cause != cause {
		t.Errorf("NewInputError() Cause = %v, want %v", err.Cause, cause)
	}
}

func TestNewFormatError(t *testing.T) {
	cause := errors.New("underlying error")
	err := NewFormatError("test message", cause)

	if err.Type != ErrorTypeFormat {
		t.Errorf("NewFormatError() Type = %v, want %v", err.Type, ErrorTypeFormat)
	}
	if err.Message != "test message" {
		t.Errorf("NewFormatError() Message = %v, want %v", err.Message, "test message")
	}
	if err.Cause != cause {
		t.Errorf("NewFormatError() Cause = %v, want %v", err.Cause, cause)
	}
}

func TestNewTransformError(t *testing.T) {
	cause := errors.New("underlying error")
	err := NewTransformError("test message", cause)

	if err.Type != ErrorTypeTransform {
		t.Errorf("NewTransformError() Type = %v, want %v", err.Type, ErrorTypeTransform)
	}
	if err.Message != "test message" {
		t.Errorf("NewTransformError() Message = %v, want %v", err.Message, "test message")
	}
	if err.Cause != cause {
		t.Errorf("NewTransformError() Cause = %v, want %v", err.Cause, cause)
	}
}

func TestNewSQLError(t *testing.T) {
	cause := errors.New("underlying error")
	err := NewSQLError("test message", cause)

	if err.Type != ErrorTypeSQL {
		t.Errorf("NewSQLError() Type = %v, want %v", err.Type, ErrorTypeSQL)
	}
	if err.Message != "test message" {
		t.Errorf("NewSQLError() Message = %v, want %v", err.Message, "test message")
	}
	if err.Cause != cause {
		t.Errorf("NewSQLError() Cause = %v, want %v", err.Cause, cause)
	}
}

func TestNewOutputError(t *testing.T) {
	cause := errors.New("underlying error")
	err := NewOutputError("test message", cause)

	if err.Type != ErrorTypeOutput {
		t.Errorf("NewOutputError() Type = %v, want %v", err.Type, ErrorTypeOutput)
	}
	if err.Message != "test message" {
		t.Errorf("NewOutputError() Message = %v, want %v", err.Message, "test message")
	}
	if err.Cause != cause {
		t.Errorf("NewOutputError() Cause = %v, want %v", err.Cause, cause)
	}
}

func TestNewInternalError(t *testing.T) {
	cause := errors.New("underlying error")
	err := NewInternalError("test message", cause)

	if err.Type != ErrorTypeInternal {
		t.Errorf("NewInternalError() Type = %v, want %v", err.Type, ErrorTypeInternal)
	}
	if err.Message != "test message" {
		t.Errorf("NewInternalError() Message = %v, want %v", err.Message, "test message")
	}
	if err.Cause != cause {
		t.Errorf("NewInternalError() Cause = %v, want %v", err.Cause, cause)
	}
}
