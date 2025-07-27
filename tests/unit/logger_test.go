package unit

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/jsas4coding/utify/pkg/logger"
	"github.com/jsas4coding/utify/pkg/messages"
)

func TestSetLogTarget(t *testing.T) {
	tempFile := "./test_utify.log"
	defer func() { _ = os.Remove(tempFile) }() // Ignore error in cleanup

	err := logger.SetLogTarget(tempFile)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if logger.GetLogTarget() != tempFile {
		t.Errorf("Expected log target %s, got %s", tempFile, logger.GetLogTarget())
	}
}

func TestLoggingEnabled(t *testing.T) {
	original := logger.IsEnabled()
	defer logger.SetEnabled(original)

	logger.SetEnabled(false)
	if logger.IsEnabled() {
		t.Error("Expected logging to be disabled")
	}

	logger.SetEnabled(true)
	if !logger.IsEnabled() {
		t.Error("Expected logging to be enabled")
	}
}

func TestLogMessage(t *testing.T) {
	tempFile := "./test_utify_message.log"
	defer func() { _ = os.Remove(tempFile) }() // Ignore error in cleanup

	err := logger.SetLogTarget(tempFile)
	if err != nil {
		t.Fatalf("Failed to set log target: %v", err)
	}
	
	logger.LogMessage(messages.Success, "Test message")
	logger.Close()

	content, err := os.ReadFile(tempFile)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	if !strings.Contains(string(content), "Test message") {
		t.Error("Log file should contain the test message")
	}

	// Test JSON structure
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
	tempFile := "./test_utify_only.log"
	defer func() { _ = os.Remove(tempFile) }() // Ignore error in cleanup

	err := logger.SetLogTarget(tempFile)
	if err != nil {
		t.Fatalf("Failed to set log target: %v", err)
	}
	
	logger.LogOnly(messages.Error, "Error only message")
	logger.Close()

	content, err := os.ReadFile(tempFile)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	if !strings.Contains(string(content), "Error only message") {
		t.Error("Log file should contain the error only message")
	}
}