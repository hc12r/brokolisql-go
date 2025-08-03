package common

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger(LogLevelInfo)

	if logger == nil {
		t.Error("NewLogger() returned nil")
	}

	if logger != nil && logger.level != LogLevelInfo {
		t.Errorf("NewLogger() level = %v, want %v", logger.level, LogLevelInfo)
	}
}

func TestNewLoggerWithWriter(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLoggerWithWriter(&buf, LogLevelDebug)

	if logger == nil {
		t.Error("NewLoggerWithWriter() returned nil")
	}

	if logger != nil && logger.level != LogLevelDebug {
		t.Errorf("NewLoggerWithWriter() level = %v, want %v", logger.level, LogLevelDebug)
	}
}

func TestLogger_SetLevel(t *testing.T) {
	logger := NewLogger(LogLevelInfo)
	logger.SetLevel(LogLevelError)

	if logger.level != LogLevelError {
		t.Errorf("SetLevel() level = %v, want %v", logger.level, LogLevelError)
	}
}

func TestLogger_Debug(t *testing.T) {
	tests := []struct {
		name      string
		level     LogLevel
		shouldLog bool
	}{
		{
			name:      "Debug level",
			level:     LogLevelDebug,
			shouldLog: true,
		},
		{
			name:      "Info level",
			level:     LogLevelInfo,
			shouldLog: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			logger := NewLoggerWithWriter(&buf, tt.level)

			logger.Debug("test message")

			if tt.shouldLog && !strings.Contains(buf.String(), "test message") {
				t.Errorf("Debug() should log message at level %v", tt.level)
			}

			if !tt.shouldLog && buf.String() != "" {
				t.Errorf("Debug() should not log message at level %v", tt.level)
			}
		})
	}
}

func TestLogger_Info(t *testing.T) {
	tests := []struct {
		name      string
		level     LogLevel
		shouldLog bool
	}{
		{
			name:      "Debug level",
			level:     LogLevelDebug,
			shouldLog: true,
		},
		{
			name:      "Info level",
			level:     LogLevelInfo,
			shouldLog: true,
		},
		{
			name:      "Warning level",
			level:     LogLevelWarning,
			shouldLog: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			logger := NewLoggerWithWriter(&buf, tt.level)

			logger.Info("test message")

			if tt.shouldLog && !strings.Contains(buf.String(), "test message") {
				t.Errorf("Info() should log message at level %v", tt.level)
			}

			if !tt.shouldLog && buf.String() != "" {
				t.Errorf("Info() should not log message at level %v", tt.level)
			}
		})
	}
}

func TestLogger_Warning(t *testing.T) {
	tests := []struct {
		name      string
		level     LogLevel
		shouldLog bool
	}{
		{
			name:      "Info level",
			level:     LogLevelInfo,
			shouldLog: true,
		},
		{
			name:      "Warning level",
			level:     LogLevelWarning,
			shouldLog: true,
		},
		{
			name:      "Error level",
			level:     LogLevelError,
			shouldLog: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			logger := NewLoggerWithWriter(&buf, tt.level)

			logger.Warning("test message")

			if tt.shouldLog && !strings.Contains(buf.String(), "test message") {
				t.Errorf("Warning() should log message at level %v", tt.level)
			}

			if !tt.shouldLog && buf.String() != "" {
				t.Errorf("Warning() should not log message at level %v", tt.level)
			}
		})
	}
}

func TestLogger_Error(t *testing.T) {
	tests := []struct {
		name      string
		level     LogLevel
		shouldLog bool
	}{
		{
			name:      "Warning level",
			level:     LogLevelWarning,
			shouldLog: true,
		},
		{
			name:      "Error level",
			level:     LogLevelError,
			shouldLog: true,
		},
		{
			name:      "Fatal level",
			level:     LogLevelFatal,
			shouldLog: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			logger := NewLoggerWithWriter(&buf, tt.level)

			logger.Error("test message")

			if tt.shouldLog && !strings.Contains(buf.String(), "test message") {
				t.Errorf("Error() should log message at level %v", tt.level)
			}

			if !tt.shouldLog && buf.String() != "" {
				t.Errorf("Error() should not log message at level %v", tt.level)
			}
		})
	}
}

// Note: We can't easily test Fatal() because it calls os.Exit(1)

func TestLogLevelFromString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  LogLevel
	}{
		{
			name:  "DEBUG",
			input: "DEBUG",
			want:  LogLevelDebug,
		},
		{
			name:  "debug (lowercase)",
			input: "debug",
			want:  LogLevelDebug,
		},
		{
			name:  "INFO",
			input: "INFO",
			want:  LogLevelInfo,
		},
		{
			name:  "WARNING",
			input: "WARNING",
			want:  LogLevelWarning,
		},
		{
			name:  "WARN",
			input: "WARN",
			want:  LogLevelWarning,
		},
		{
			name:  "ERROR",
			input: "ERROR",
			want:  LogLevelError,
		},
		{
			name:  "FATAL",
			input: "FATAL",
			want:  LogLevelFatal,
		},
		{
			name:  "Invalid",
			input: "INVALID",
			want:  LogLevelInfo, // Default
		},
		{
			name:  "Empty",
			input: "",
			want:  LogLevelInfo, // Default
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LogLevelFromString(tt.input); got != tt.want {
				t.Errorf("LogLevelFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultLogger(t *testing.T) {
	if DefaultLogger == nil {
		t.Error("DefaultLogger is nil")
	}

	if DefaultLogger.level != LogLevelInfo {
		t.Errorf("DefaultLogger level = %v, want %v", DefaultLogger.level, LogLevelInfo)
	}
}
