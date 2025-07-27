package unit

import (
	"errors"
	"strings"
	"testing"

	testutil "github.com/jsas4coding/utify/internal/tests"
	"github.com/jsas4coding/utify/pkg/colors"
	"github.com/jsas4coding/utify/pkg/formatter"
	"github.com/jsas4coding/utify/pkg/messages"
	"github.com/jsas4coding/utify/pkg/options"
)

func TestEcho(t *testing.T) {
	tests := []struct {
		name        string
		msgType     messages.Type
		text        string
		opts        *options.Options
		shouldError bool
	}{
		{"Success", messages.Success, "Operation completed", options.Default(), false},
		{"Error", messages.Error, "An error occurred", options.Default(), true},
		{"Warning", messages.Warning, "This is a warning", options.Default(), false},
		{"Info", messages.Info, "Just info", options.Default(), false},
		{"Debug", messages.Debug, "Debugging", options.Default(), true},
		{"Critical", messages.Critical, "Critical!", options.Default(), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := testutil.CaptureOutput(func() {
				_, _ = formatter.Echo(tt.msgType, tt.text, tt.opts)
			})

			if !strings.Contains(output, tt.text) {
				t.Errorf("Expected output to contain %q, got %q", tt.text, output)
			}

			_, err := formatter.Echo(tt.msgType, tt.text, tt.opts)
			if tt.shouldError && !errors.Is(err, formatter.ErrSilent) {
				t.Errorf("Expected ErrSilent for %s, got %v", tt.name, err)
			} else if !tt.shouldError && err != nil {
				t.Errorf("Expected nil error for %s, got %v", tt.name, err)
			}
		})
	}
}

func TestEchoBold(t *testing.T) {
	opts := options.Default().WithBold()
	
	output := testutil.CaptureOutput(func() {
		_, _ = formatter.Echo(messages.Success, "Bold text", opts)
	})

	if !strings.Contains(output, colors.Bold) {
		t.Errorf("Expected output to contain bold formatting %q, got %q", colors.Bold, output)
	}
}

func TestEchoItalic(t *testing.T) {
	opts := options.Default().WithItalic()
	
	output := testutil.CaptureOutput(func() {
		_, _ = formatter.Echo(messages.Success, "Italic text", opts)
	})

	if !strings.Contains(output, colors.Italic) {
		t.Errorf("Expected output to contain italic formatting %q, got %q", colors.Italic, output)
	}
}

func TestEchoNoColor(t *testing.T) {
	opts := options.Default().WithoutColor()
	
	output := testutil.CaptureOutput(func() {
		_, _ = formatter.Echo(messages.Success, "Plain text", opts)
	})

	if strings.Contains(output, colors.Green) {
		t.Errorf("Expected output to not contain color formatting, got %q", output)
	}
}

func TestEchoNoStyle(t *testing.T) {
	opts := options.Default().WithBold().WithItalic().WithoutStyle()
	
	output := testutil.CaptureOutput(func() {
		_, _ = formatter.Echo(messages.Success, "No style", opts)
	})

	if strings.Contains(output, colors.Bold) || strings.Contains(output, colors.Italic) {
		t.Errorf("Expected output to not contain style formatting, got %q", output)
	}
}

func TestEchoCallback(t *testing.T) {
	var callbackType messages.Type
	var callbackText string
	
	callback := func(msgType messages.Type, text string) {
		callbackType = msgType
		callbackText = text
	}
	
	opts := options.Default().WithCallback(callback)
	
	_, _ = formatter.Echo(messages.Success, "Testing callback", opts)
	
	if callbackType != messages.Success {
		t.Errorf("Expected callback type to be %v, got %v", messages.Success, callbackType)
	}
	
	if callbackText != "Testing callback" {
		t.Errorf("Expected callback text to be %q, got %q", "Testing callback", callbackText)
	}
}

func TestEchoLogsToFile(t *testing.T) {
	// This test just verifies that Echo calls the logger without errors
	// The actual logging functionality is tested in logger_test.go
	output := testutil.CaptureOutput(func() {
		_, _ = formatter.Echo(messages.Success, "Log test", options.Default())
	})

	if !strings.Contains(output, "Log test") {
		t.Errorf("Expected output to contain %q, got %q", "Log test", output)
	}
}