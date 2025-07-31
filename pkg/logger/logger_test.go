package logger

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jsas4coding/utify/internal/tests"
	"github.com/jsas4coding/utify/pkg/messages"
)

func TestSetLogTarget(t *testing.T) {
	tests.CreateDataDir(t)
	tempFile := filepath.Join(tests.DataDir, "test_utify.log")
	defer func() { _ = os.Remove(tempFile) }()

	err := SetLogTarget(tempFile)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if GetLogTarget() != tempFile {
		t.Errorf("Expected log target %s, got %s", tempFile, GetLogTarget())
	}
}

func TestLoggingEnabled(t *testing.T) {
	original := IsEnabled()
	defer SetEnabled(original)

	SetEnabled(false)
	if IsEnabled() {
		t.Error("Expected logging to be disabled")
	}

	SetEnabled(true)
	if !IsEnabled() {
		t.Error("Expected logging to be enabled")
	}
}

func TestLogMessage(t *testing.T) {
	tests.CreateDataDir(t)
	tempFile := filepath.Join(tests.DataDir, "test_utify_message.log")
	defer func() { _ = os.Remove(tempFile) }()

	err := SetLogTarget(tempFile)
	if err != nil {
		t.Fatalf("Failed to set log target: %v", err)
	}

	LogMessage(messages.Success, "Test message")
	Close()

	content, err := os.ReadFile(tempFile)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	if !strings.Contains(string(content), "Test message") {
		t.Error("Log file should contain the test message")
	}

	var logEntry map[string]interface{}
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	if len(lines) > 0 {
		err = json.Unmarshal([]byte(lines[0]), &logEntry)
		if err != nil {
			t.Errorf("Log entry should be valid JSON: %v", err)
		}

		if logEntry["message"] != "Test message" {
			t.Errorf("Expected message 'Test message', got %v", logEntry["message"])
		}

		if logEntry["level"] != "SUCCESS" {
			t.Errorf("Expected level 'SUCCESS', got %v", logEntry["level"])
		}

		if logEntry["type"] != "success" {
			t.Errorf("Expected type 'success', got %v", logEntry["type"])
		}
	}
}

func TestLogOnly(t *testing.T) {
	tests.CreateDataDir(t)
	tempFile := filepath.Join(tests.DataDir, "test_utify_only.log")
	defer func() { _ = os.Remove(tempFile) }()

	err := SetLogTarget(tempFile)
	if err != nil {
		t.Fatalf("Failed to set log target: %v", err)
	}

	LogOnly(messages.Error, "Error only message")
	Close()

	content, err := os.ReadFile(tempFile)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	if !strings.Contains(string(content), "Error only message") {
		t.Error("Log file should contain the error only message")
	}
}

func TestLogFilePermissions(t *testing.T) {
	tests.CreateDataDir(t)
	tempFile := filepath.Join(tests.DataDir, "test_permissions.log")
	defer func() { _ = os.Remove(tempFile) }()

	err := SetLogTarget(tempFile)
	if err != nil {
		t.Fatalf("Failed to set log target: %v", err)
	}
	Close()

	info, err := os.Stat(tempFile)
	if err != nil {
		t.Fatalf("Failed to stat log file: %v", err)
	}

	if info.Mode().Perm() != 0644 {
		t.Errorf("Expected file permissions 0644, got %v", info.Mode().Perm())
	}
}

func TestLogToInvalidTarget(t *testing.T) {
	// Attempt to log to a directory that doesn't exist and we can't create
	invalidTarget := "/nonexistent/dir/test.log"
	err := SetLogTarget(invalidTarget)
	if err == nil {
		t.Error("Expected an error when setting an invalid log target, but got nil")
	}
}

func TestLogMessageWithDisabledLogging(t *testing.T) {
	tests.CreateDataDir(t)
	tempFile := filepath.Join(tests.DataDir, "test_disabled_logging.log")
	defer func() { _ = os.Remove(tempFile) }()

	err := SetLogTarget(tempFile)
	if err != nil {
		t.Fatalf("Failed to set log target: %v", err)
	}

	// Disable logging
	SetEnabled(false)
	defer SetEnabled(true)

	LogMessage(messages.Success, "Disabled logging message")
	Close()

	content, err := os.ReadFile(tempFile)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	if strings.Contains(string(content), "Disabled logging message") {
		t.Error("Disabled logging should not write to log file")
	}
}

func TestGetBinaryName(t *testing.T) {
	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	os.Args = []string{"/path/to/some/binary"}
	name := getBinaryName()
	if name != "binary" {
		t.Errorf("Expected binary name 'binary', got %q", name)
	}

	os.Args = []string{}
	name = getBinaryName()
	if name != "utify" {
		t.Errorf("Expected fallback binary name 'utify', got %q", name)
	}
}