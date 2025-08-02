package errors

import (
	"bytes"
	stderrors "errors"
	"log"
	"strings"
	"testing"
)

func TestCheckError(t *testing.T) {
	// Save the original log settings and restore them after the test
	originalOutput := log.Writer()
	originalFlags := log.Flags()
	originalPrefix := log.Prefix()

	// Create a buffer to capture log output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetFlags(0) // Remove timestamp from the output
	log.SetPrefix("")

	// Create backups of the original functions
	originalOsExit := osExit
	originalLogFatalf := logFatalf

	defer func() {
		// Restore original functions and settings
		log.SetOutput(originalOutput)
		log.SetFlags(originalFlags)
		log.SetPrefix(originalPrefix)
		osExit = originalOsExit
		logFatalf = originalLogFatalf
	}()

	// Override functions with test versions
	exitCalled := false
	osExit = func(code int) {
		exitCalled = true
	}

	logFatalf = func(format string, v ...interface{}) {
		log.Printf(format, v...)
		osExit(1)
	}

	t.Run("nil error should not call log.Fatalf", func(t *testing.T) {
		// Reset the buffer and exitCalled flag
		buf.Reset()
		exitCalled = false

		// Call the function with nil error
		CheckError(nil)

		// Check that log.Fatalf was not called
		if buf.Len() > 0 {
			t.Errorf("Expected no log output, got: %s", buf.String())
		}

		// Check that os.Exit was not called
		if exitCalled {
			t.Errorf("os.Exit was called unexpectedly")
		}
	})

	t.Run("non-nil error should call log.Fatalf", func(t *testing.T) {
		// Reset the buffer and exitCalled flag
		buf.Reset()
		exitCalled = false

		// Create a test error
		testErr := stderrors.New("test error")

		// Call the function with non-nil error
		// This will call our mocked osExit function
		CheckError(testErr)

		// Check that log.Fatalf was called with the expected message
		logOutput := buf.String()
		if !strings.Contains(logOutput, "error: test error") {
			t.Errorf("Expected log output to contain 'error: test error', got: %s", logOutput)
		}

		// Check that os.Exit was called
		if !exitCalled {
			t.Errorf("os.Exit was not called")
		}
	})
}
