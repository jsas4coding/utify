package integration

import (
	"os"
	"strings"
	"testing"

	"github.com/jsas4coding/utify"
	testutil "github.com/jsas4coding/utify/internal/tests"
)

func TestPublicAPICompatibility(t *testing.T) {
	tests := []struct {
		name string
		fn   func()
	}{
		{
			"Success",
			func() { utify.Success("Test success", utify.OptionsDefault()) },
		},
		{
			"Error",
			func() { utify.Error("Test error", utify.OptionsDefault()) },
		},
		{
			"Warning",
			func() { utify.Warning("Test warning", utify.OptionsDefault()) },
		},
		{
			"Info",
			func() { utify.Info("Test info", utify.OptionsDefault()) },
		},
		{
			"Debug",
			func() { utify.Debug("Test debug", utify.OptionsDefault()) },
		},
		{
			"Critical",
			func() { utify.Critical("Test critical", utify.OptionsDefault()) },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := testutil.CaptureOutput(tt.fn)
			
			if !strings.Contains(output, "Test") {
				t.Errorf("Expected output to contain 'Test', got %q", output)
			}
		})
	}
}

func TestFormattedFunctions(t *testing.T) {
	tests := []struct {
		name string
		fn   func()
	}{
		{
			"Successf",
			func() { utify.Successf("Test %s", utify.OptionsDefault(), "success") },
		},
		{
			"Errorf",
			func() { utify.Errorf("Test %s", utify.OptionsDefault(), "error") },
		},
		{
			"Warningf",
			func() { utify.Warningf("Test %s", utify.OptionsDefault(), "warning") },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := testutil.CaptureOutput(tt.fn)
			
			if !strings.Contains(output, "Test") {
				t.Errorf("Expected output to contain 'Test', got %q", output)
			}
		})
	}
}

func TestGetFunctions(t *testing.T) {
	text, err := utify.GetSuccess("Test success", utify.OptionsDefault())
	if text != "Test success" {
		t.Errorf("Expected text to be 'Test success', got %q", text)
	}
	if err != nil {
		t.Errorf("Expected no error for success, got %v", err)
	}

	text, err = utify.GetError("Test error", utify.OptionsDefault())
	if text != "Test error" {
		t.Errorf("Expected text to be 'Test error', got %q", text)
	}
	if err != utify.ErrSilent {
		t.Errorf("Expected ErrSilent for error, got %v", err)
	}
}

func TestOptionsChaining(t *testing.T) {
	opts := utify.OptionsDefault().
		WithBold().
		WithItalic().
		WithoutColor()

	if !opts.Bold {
		t.Error("Expected Bold to be true")
	}
	if !opts.Italic {
		t.Error("Expected Italic to be true")
	}
	if !opts.NoColor {
		t.Error("Expected NoColor to be true")
	}
}

func TestColorTableOverride(t *testing.T) {
	customColor := "\033[95m"
	utify.SetColorTable(map[string]string{
		string(utify.MessageSuccess): customColor,
	})

	output := testutil.CaptureOutput(func() {
		utify.Success("Custom color test", utify.OptionsDefault().WithoutStyle())
	})

	if !strings.Contains(output, customColor) {
		t.Errorf("Expected output to contain custom color %q, got %q", customColor, output)
	}
}

func TestLoggingFunctionality(t *testing.T) {
	// Test log target configuration
	original := utify.GetLogTarget()
	defer utify.SetLogTarget(original)

	tempTarget := "./test_integration.log"
	defer func() {
		utify.CloseLogger()
		// Clean up log files
		_ = os.Remove(tempTarget)
		_ = os.Remove("./test_integration.log.test")
	}()

	err := utify.SetLogTarget(tempTarget)
	if err != nil {
		t.Errorf("Failed to set log target: %v", err)
	}

	if utify.GetLogTarget() != tempTarget {
		t.Errorf("Expected log target %s, got %s", tempTarget, utify.GetLogTarget())
	}

	// Test logging enabled/disabled
	if !utify.IsLoggingEnabled() {
		t.Error("Expected logging to be enabled by default")
	}

	utify.SetLoggingEnabled(false)
	if utify.IsLoggingEnabled() {
		t.Error("Expected logging to be disabled")
	}

	utify.SetLoggingEnabled(true)
	if !utify.IsLoggingEnabled() {
		t.Error("Expected logging to be re-enabled")
	}
}

func TestLogOnlyFunctions(t *testing.T) {
	tempTarget := "./test_log_only.log"
	defer func() {
		utify.CloseLogger()
		os.Remove(tempTarget)
	}()

	err := utify.SetLogTarget(tempTarget)
	if err != nil {
		t.Errorf("Failed to set log target: %v", err)
	}

	// Test log-only functions (should not output to stdout)
	output := testutil.CaptureOutput(func() {
		utify.LogInfo("This should only be logged")
		utify.LogSuccessf("Successfully processed %d items", 5)
	})

	if strings.Contains(output, "This should only be logged") {
		t.Error("LogInfo should not output to stdout")
	}

	if strings.Contains(output, "Successfully processed") {
		t.Error("LogSuccessf should not output to stdout")
	}
}

func TestIconFunctionality(t *testing.T) {
	// Test icon detection
	detected := utify.IsNerdFontDetected()
	_ = detected // Just ensure it works

	// Test icon type functions
	originalType := utify.GetIconType()
	defer func() {
		// Reset to original type
		switch originalType {
		case 0:
			utify.DisableIcons()
		case 1:
			utify.ForceRegularIcons()
		case 2:
			utify.ForceNerdFont()
		}
	}()

	// Test with icons enabled
	utify.ForceRegularIcons()
	output := testutil.CaptureOutput(func() {
		utify.Success("Test with icon", utify.OptionsDefault().WithIcon())
	})

	if !strings.Contains(output, "✅") {
		t.Error("Expected success icon ✅ in output")
	}

	// Test with icons disabled
	output = testutil.CaptureOutput(func() {
		utify.Success("Test without icon", utify.OptionsDefault().WithoutIcon())
	})

	if strings.Contains(output, "✅") {
		t.Error("Should not contain icon when WithoutIcon() is used")
	}

	// Test Nerd Font icons
	utify.ForceNerdFont()
	output = testutil.CaptureOutput(func() {
		utify.Error("Nerd font test", utify.OptionsDefault().WithIcon())
	})

	// Should contain some icon (Nerd Font icons are harder to test visually)
	if !strings.Contains(output, "Nerd font test") {
		t.Error("Should contain the message text")
	}
}