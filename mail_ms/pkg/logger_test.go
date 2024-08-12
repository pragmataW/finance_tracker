package pkg

import (
	"bytes"
	"testing"
)

func TestNewLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(INFO, &buf)

	if logger == nil {
		t.Fatal("Expected logger to be initialized, got nil")
	}

	// Check if the logger's level is set correctly
	if logger.logLevel != INFO {
		t.Fatalf("Expected log level %v, got %v", INFO, logger.logLevel)
	}
}

func TestInfoLogging(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(INFO, &buf)

	logger.Info("Info message")
	expected := "INFO: "
	if !bytes.HasPrefix(buf.Bytes(), []byte(expected)) {
		t.Fatalf("Expected log to start with %q, got %q", expected, buf.String())
	}
	if !bytes.Contains(buf.Bytes(), []byte("Info message")) {
		t.Fatalf("Expected log to contain %q, got %q", "Info message", buf.String())
	}
}

func TestWarningLogging(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(WARNING, &buf)

	logger.Warning("Warning message")
	expected := "WARNING: "
	if !bytes.HasPrefix(buf.Bytes(), []byte(expected)) {
		t.Fatalf("Expected log to start with %q, got %q", expected, buf.String())
	}
	if !bytes.Contains(buf.Bytes(), []byte("Warning message")) {
		t.Fatalf("Expected log to contain %q, got %q", "Warning message", buf.String())
	}
}

func TestErrorLogging(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(ERROR, &buf)

	logger.Error("Error message")
	expected := "ERROR: "
	if !bytes.HasPrefix(buf.Bytes(), []byte(expected)) {
		t.Fatalf("Expected log to start with %q, got %q", expected, buf.String())
	}
	if !bytes.Contains(buf.Bytes(), []byte("Error message")) {
		t.Fatalf("Expected log to contain %q, got %q", "Error message", buf.String())
	}
}

func TestDebugLogging(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(DEBUG, &buf)

	logger.Debug("Debug message")
	expected := "DEBUG: "
	if !bytes.HasPrefix(buf.Bytes(), []byte(expected)) {
		t.Fatalf("Expected log to start with %q, got %q", expected, buf.String())
	}
	if !bytes.Contains(buf.Bytes(), []byte("Debug message")) {
		t.Fatalf("Expected log to contain %q, got %q", "Debug message", buf.String())
	}
}

func TestLogLevelFiltering(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(INFO, &buf)

	logger.Debug("Debug message")
	if buf.Len() > 0 {
		t.Fatal("Expected no log output for DEBUG level when log level is INFO")
	}
}