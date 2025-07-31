package messages

import (
	"testing"

	"github.com/jsas4coding/utify/pkg/colors"
)

func TestGetColor(t *testing.T) {
	tests := []struct {
		name     string
		msgType  Type
		expected string
	}{
		{"Success", Success, colors.Green},
		{"Error", Error, colors.Red},
		{"Warning", Warning, colors.Yellow},
		{"Info", Info, colors.Cyan},
		{"Debug", Debug, colors.Gray},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			color := GetColor(tt.msgType)
			if color != tt.expected {
				t.Errorf("Expected color %q for %s, got %q", tt.expected, tt.msgType, color)
			}
		})
	}
}

func TestIsErrorType(t *testing.T) {
	tests := []struct {
		name    string
		msgType Type
		isError bool
	}{
		{"Error", Error, true},
		{"Critical", Critical, true},
		{"Debug", Debug, true},
		{"Success", Success, false},
		{"Warning", Warning, false},
		{"Info", Info, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsErrorType(tt.msgType)
			if result != tt.isError {
				t.Errorf("Expected IsErrorType(%s) to be %v, got %v", tt.msgType, tt.isError, result)
			}
		})
	}
}

func TestGetColorWithUserOverride(t *testing.T) {
	colors.ClearUserColors()
	customColor := "\033[95m"
	colors.SetColorTable(map[string]string{
		string(Success): customColor,
	})

	color := GetColor(Success)
	if color != customColor {
		t.Errorf("Expected custom color %q for success, got %q", customColor, color)
	}

	// Clean up
	colors.ClearUserColors()
}
