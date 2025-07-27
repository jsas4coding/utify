package unit

import (
	"testing"

	"github.com/jsas4coding/utify/pkg/colors"
)

func TestSetColorTable(t *testing.T) {
	colors.ClearUserColors()
	
	customColors := map[string]string{
		"success": "\033[95m", // Purple
		"error":   "\033[96m", // Cyan
	}
	
	colors.SetColorTable(customColors)
	
	color, exists := colors.GetUserColor("success")
	if !exists {
		t.Error("Expected success color to exist in user colors")
	}
	if color != "\033[95m" {
		t.Errorf("Expected success color to be %q, got %q", "\033[95m", color)
	}
	
	color, exists = colors.GetUserColor("error")
	if !exists {
		t.Error("Expected error color to exist in user colors")
	}
	if color != "\033[96m" {
		t.Errorf("Expected error color to be %q, got %q", "\033[96m", color)
	}
}

func TestGetUserColorNotExists(t *testing.T) {
	colors.ClearUserColors()
	
	_, exists := colors.GetUserColor("nonexistent")
	if exists {
		t.Error("Expected nonexistent color to not exist")
	}
}

func TestClearUserColors(t *testing.T) {
	colors.SetColorTable(map[string]string{"test": "value"})
	colors.ClearUserColors()
	
	_, exists := colors.GetUserColor("test")
	if exists {
		t.Error("Expected user colors to be cleared")
	}
}