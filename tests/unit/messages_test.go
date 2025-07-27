package unit

import (
	"testing"

	"github.com/jsas4coding/utify/pkg/colors"
	"github.com/jsas4coding/utify/pkg/messages"
)

func TestGetColor(t *testing.T) {
	tests := []struct {
		name     string
		msgType  messages.Type
		expected string
	}{
		{"Success", messages.Success, colors.Green},
		{"Error", messages.Error, colors.Red},
		{"Warning", messages.Warning, colors.Yellow},
		{"Info", messages.Info, colors.Cyan},
		{"Debug", messages.Debug, colors.Gray},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			color := messages.GetColor(tt.msgType)
			if color != tt.expected {
				t.Errorf("Expected color %q for %s, got %q", tt.expected, tt.msgType, color)
			}
		})
	}
}

func TestIsErrorType(t *testing.T) {
	tests := []struct {
		name    string
		msgType messages.Type
		isError bool
	}{
		{"Error", messages.Error, true},
		{"Critical", messages.Critical, true},
		{"Debug", messages.Debug, true},
		{"Success", messages.Success, false},
		{"Warning", messages.Warning, false},
		{"Info", messages.Info, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := messages.IsErrorType(tt.msgType)
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
		string(messages.Success): customColor,
	})

	color := messages.GetColor(messages.Success)
	if color != customColor {
		t.Errorf("Expected custom color %q for success, got %q", customColor, color)
	}

	// Clean up
	colors.ClearUserColors()
}
