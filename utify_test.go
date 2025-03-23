package utify

import (
	"bytes"
	"errors"
	"log"
	"os"
	"strings"
	"testing"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old
	_, _ = buf.ReadFrom(r)
	return buf.String()
}

func TestEcho(t *testing.T) {
	tests := []struct {
		name     string
		msgType  MessageType
		text     string
		options  *Options
		expected string
	}{
		{"Success", MessageSuccess, "Operation completed", OptionsDefault(), "Operation completed"},
		{"Error", MessageError, "An error occurred", OptionsDefault(), "An error occurred"},
		{"Warning", MessageWarning, "This is a warning", OptionsDefault(), "This is a warning"},
		{"Info", MessageInfo, "Just info", OptionsDefault(), "Just info"},
		{"Debug", MessageDebug, "Debugging", OptionsDefault(), "Debugging"},
		{"Critical", MessageCritical, "Critical!", OptionsDefault(), "Critical!"},
		{"Bold", MessageSuccess, "Bold text", OptionsDefault().WithBold(), StyleBold},
		{"Italic", MessageSuccess, "Italic text", OptionsDefault().WithItalic(), StyleItalic},
		{"NoColor", MessageSuccess, "Plain text", OptionsDefault().WithoutColor(), "Plain text"},
		{"NoStyle", MessageSuccess, "No style", OptionsDefault().WithoutStyle(), "No style"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				_, _ = Echo(tt.msgType, tt.text, tt.options)
			})

			if !strings.Contains(output, tt.expected) {
				t.Errorf("Expected output to contain %q, got %q", tt.expected, output)
			}
		})
	}
}

func TestMessageFunctions(t *testing.T) {
	tests := []struct {
		name    string
		fn      func(string, *Options)
		message string
		isError bool
	}{
		{"Success", Success, "Success!", false},
		{"Error", Error, "Something failed", true},
		{"Warning", Warning, "Watch out!", false},
		{"Info", Info, "FYI", false},
		{"Debug", Debug, "Debugging", true},
		{"Critical", Critical, "Boom!", true},
		{"Delete", Delete, "Removed", false},
		{"Update", Update, "Changed", false},
		{"Install", Install, "Setup done", false},
		{"Upgrade", Upgrade, "Upgraded", false},
		{"Edit", Edit, "Modified", false},
		{"New", New, "Created", false},
		{"Download", Download, "Got it", false},
		{"Upload", Upload, "Sent", false},
		{"Sync", Sync, "Synchronized", false},
		{"Search", Search, "Found something", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				tt.fn(tt.message, OptionsDefault())
			})

			if !strings.Contains(output, tt.message) {
				t.Errorf("Expected output to contain %q, got %q", tt.message, output)
			}
		})
	}
}

func TestExitDisablesCallback(t *testing.T) {
	opts := OptionsDefault().WithExit()

	if opts.Callback != nil {
		t.Errorf("Callback should be nil when Exit is enabled")
	}
	if !opts.Exit {
		t.Errorf("Exit should be true")
	}
}

func TestCallbackDisablesExit(t *testing.T) {
	var called bool
	cb := func(msgType MessageType, text string) {
		called = true
	}

	opts := OptionsDefault().WithCallback(cb)
	_, _ = Echo(MessageSuccess, "Testing callback", opts)

	if !called {
		t.Errorf("Expected callback to be called")
	}
	if opts.Exit {
		t.Errorf("Expected Exit to be disabled when Callback is set")
	}
}

func TestFormattedFunctions(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string, *Options, ...any)
	}{
		{"Successf", Successf},
		{"Errorf", Errorf},
		{"Warningf", Warningf},
		{"Infof", Infof},
		{"Debugf", Debugf},
		{"Criticalf", Criticalf},
		{"Deletef", Deletef},
		{"Updatef", Updatef},
		{"Installf", Installf},
		{"Upgradef", Upgradef},
		{"Editf", Editf},
		{"Newf", Newf},
		{"Downloadf", Downloadf},
		{"Uploadf", Uploadf},
		{"Syncf", Syncf},
		{"Searchf", Searchf},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := "Formatted " + tt.name
			output := captureOutput(func() {
				tt.fn("Formatted %s", OptionsDefault(), tt.name)
			})

			if !strings.Contains(output, expected) {
				t.Errorf("Expected output to contain %q, got %q", expected, output)
			}
		})
	}
}

func TestGetFunctions(t *testing.T) {
	tests := []struct {
		name    string
		fn      func(string, *Options) (string, error)
		message string
		isError bool
	}{
		{"GetSuccess", GetSuccess, "Success!", false},
		{"GetError", GetError, "Something failed", true},
		{"GetWarning", GetWarning, "Watch out!", false},
		{"GetInfo", GetInfo, "FYI", false},
		{"GetDebug", GetDebug, "Debugging", true},
		{"GetCritical", GetCritical, "Boom!", true},
		{"GetDelete", GetDelete, "Removed", false},
		{"GetUpdate", GetUpdate, "Changed", false},
		{"GetInstall", GetInstall, "Setup done", false},
		{"GetUpgrade", GetUpgrade, "Upgraded", false},
		{"GetEdit", GetEdit, "Modified", false},
		{"GetNew", GetNew, "Created", false},
		{"GetDownload", GetDownload, "Got it", false},
		{"GetUpload", GetUpload, "Sent", false},
		{"GetSync", GetSync, "Synchronized", false},
		{"GetSearch", GetSearch, "Found something", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				_, _ = tt.fn(tt.message, OptionsDefault())
			})

			if !strings.Contains(output, tt.message) {
				t.Errorf("Expected output to contain %q, got %q", tt.message, output)
			}

			_, err := tt.fn(tt.message, OptionsDefault())
			if tt.isError && !errors.Is(err, ErrSilent) {
				t.Errorf("Expected ErrSilent for %s, got %v", tt.name, err)
			} else if !tt.isError && err != nil {
				t.Errorf("Expected nil error for %s, got %v", tt.name, err)
			}
		})
	}
}

func TestGetFormattedFunctions(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string, *Options, ...any) (string, error)
	}{
		{"GetSuccessf", GetSuccessf},
		{"GetErrorf", GetErrorf},
		{"GetWarningf", GetWarningf},
		{"GetInfof", GetInfof},
		{"GetDebugf", GetDebugf},
		{"GetCriticalf", GetCriticalf},
		{"GetDeletef", GetDeletef},
		{"GetUpdatef", GetUpdatef},
		{"GetInstallf", GetInstallf},
		{"GetUpgradef", GetUpgradef},
		{"GetEditf", GetEditf},
		{"GetNewf", GetNewf},
		{"GetDownloadf", GetDownloadf},
		{"GetUploadf", GetUploadf},
		{"GetSyncf", GetSyncf},
		{"GetSearchf", GetSearchf},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := "Formatted " + tt.name
			output := captureOutput(func() {
				_, _ = tt.fn("Formatted %s", OptionsDefault(), tt.name)
			})

			if !strings.Contains(output, expected) {
				t.Errorf("Expected output to contain %q, got %q", expected, output)
			}
		})
	}
}

func TestSetColorTableOverride(t *testing.T) {
	customColor := "\033[95m" // Purple
	SetColorTable(map[string]string{
		string(MessageSuccess): customColor,
	})

	opts := OptionsDefault().WithoutStyle()
	output := captureOutput(func() {
		Success("Custom success color", opts)
	})

	if !strings.Contains(output, customColor) {
		t.Errorf("Expected output to contain custom color code %q, got %q", customColor, output)
	}
}

func captureLogOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

	f()

	return buf.String()
}

func TestLogOutput(t *testing.T) {
	tests := []struct {
		name    string
		fn      func(string, *Options)
		msgType MessageType
	}{
		{"SuccessLog", Success, MessageSuccess},
		{"ErrorLog", Error, MessageError},
		{"DebugLog", Debug, MessageDebug},
		{"CriticalLog", Critical, MessageCritical},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := "Log test message"
			logOut := captureLogOutput(func() {
				tt.fn(msg, OptionsDefault())
			})

			expected := "[" + strings.ToUpper(string(tt.msgType)) + "] " + msg
			if !strings.Contains(logOut, expected) {
				t.Errorf("Expected log output to contain %q, got %q", expected, logOut)
			}
		})
	}
}
