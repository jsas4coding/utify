package colors

import (
	"testing"
)

func TestSetColorTable(t *testing.T) {
	ClearUserColors()

	customColors := map[string]string{
		"success": "\033[95m", // Purple
		"error":   "\033[96m", // Cyan
	}

	SetColorTable(customColors)

	color, exists := GetUserColor("success")
	if !exists {
		t.Error("Expected success color to exist in user colors")
	}
	if color != "\033[95m" {
		t.Errorf("Expected success color to be %q, got %q", "\033[95m", color)
	}

	color, exists = GetUserColor("error")
	if !exists {
		t.Error("Expected error color to exist in user colors")
	}
	if color != "\033[96m" {
		t.Errorf("Expected error color to be %q, got %q", "\033[96m", color)
	}
}

func TestGetUserColorNotExists(t *testing.T) {
	ClearUserColors()

	_, exists := GetUserColor("nonexistent")
	if exists {
		t.Error("Expected nonexistent color to not exist")
	}
}

func TestClearUserColors(t *testing.T) {
	SetColorTable(map[string]string{"test": "value"})
	ClearUserColors()

	_, exists := GetUserColor("test")
	if exists {
		t.Error("Expected user colors to be cleared")
	}
}
