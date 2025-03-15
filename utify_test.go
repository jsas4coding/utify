package utify

import (
	"bytes"
	"os"
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
	buf.ReadFrom(r)
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
		{"Success Message", MessageSuccess, "Operation completed", OptionsDefault(), "Operation completed"},
		{"Error Message", MessageError, "An error occurred", OptionsDefault(), "An error occurred"},
		{"Warning Message", MessageWarning, "This is a warning", OptionsDefault(), "This is a warning"},
		{"Info Message", MessageInfo, "Information", OptionsDefault(), "Information"},
		{"Debug Message", MessageDebug, "Debug mode", OptionsDefault(), "Debug mode"},
		{"Critical Message", MessageCritical, "Critical failure", OptionsDefault(), "Critical failure"},
		{"Bold Option", MessageSuccess, "Bold text", OptionsDefault().WithBold(), "\033[1m"},
		{"Italic Option", MessageSuccess, "Italic text", OptionsDefault().WithItalic(), "\033[3m"},
		{"NoColor Option", MessageSuccess, "No color", OptionsDefault().WithoutColor(), ""},
		{"NoStyle Option", MessageSuccess, "No style", OptionsDefault().WithoutStyle(), ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				Echo(tt.msgType, tt.text, tt.options)
			})

			if !bytes.Contains([]byte(output), []byte(tt.expected)) {
				t.Errorf("Expected output to contain: %q, but got: %q", tt.expected, output)
			}
		})
	}
}

func TestExitDisablesCallback(t *testing.T) {
	opts := OptionsDefault().WithExit().WithCallback(func(msgType MessageType, text string) {
		t.Errorf("Callback should not be executed when Exit is enabled")
	})

	if opts.Callback != nil {
		t.Errorf("Callback should be nil when Exit is enabled")
	}
}

func TestCallbackDisablesExit(t *testing.T) {
	var callbackExecuted bool
	callback := func(msgType MessageType, text string) {
		callbackExecuted = true
	}

	opts := OptionsDefault().WithCallback(callback).WithExit()

	if opts.Exit {
		t.Errorf("Exit should be disabled when Callback is set")
	}

	Echo(MessageSuccess, "Test message", opts)

	if !callbackExecuted {
		t.Errorf("Callback should have been executed")
	}
}

func TestAllMessageFunctions(t *testing.T) {
	tests := []struct {
		name    string
		fn      func(string, *Options)
		message string
	}{
		{"Success", Success, "Success message"},
		{"Error", Error, "Error message"},
		{"Warning", Warning, "Warning message"},
		{"Info", Info, "Info message"},
		{"Debug", Debug, "Debug message"},
		{"Critical", Critical, "Critical message"},
		{"Delete", Delete, "Delete message"},
		{"Update", Update, "Update message"},
		{"Install", Install, "Install message"},
		{"Upgrade", Upgrade, "Upgrade message"},
		{"Edit", Edit, "Edit message"},
		{"New", New, "New message"},
		{"Download", Download, "Download message"},
		{"Upload", Upload, "Upload message"},
		{"Sync", Sync, "Sync message"},
		{"Search", Search, "Search message"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				tt.fn(tt.message, OptionsDefault())
			})

			if !bytes.Contains([]byte(output), []byte(tt.message)) {
				t.Errorf("Expected output to contain: %q, but got: %q", tt.message, output)
			}
		})
	}
}
