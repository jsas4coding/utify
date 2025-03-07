package utify

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer

	log.SetOutput(&buf)

	defer log.SetOutput(os.Stderr)

	f()

	return buf.String()
}

func TestEcho(t *testing.T) {
	output := captureOutput(func() {
		Echo(MessageSuccess, "Default Echo Test", Options{})
	})

	if output == "" {
		t.Errorf("Expected default Echo output, got empty string")
	}

	output = captureOutput(func() {
		Echo(MessageSuccess, "Bold Echo Test", Options{Bold: true})
	})

	if output == "" {
		t.Errorf("Expected bold Echo output, got empty string")
	}

	output = captureOutput(func() {
		Echo(MessageSuccess, "Italic Echo Test", Options{Italic: true})
	})

	if output == "" {
		t.Errorf("Expected italic Echo output, got empty string")
	}

	output = captureOutput(func() {
		Echo(MessageSuccess, "NoColor Echo Test", Options{NoColor: true})
	})

	if output == "" {
		t.Errorf("Expected no color Echo output, got empty string")
	}

	output = captureOutput(func() {
		Echo(MessageSuccess, "NoIcon Echo Test", Options{NoIcon: true})
	})

	if output == "" {
		t.Errorf("Expected no icon Echo output, got empty string")
	}

	if os.Getenv("TEST_ECHO_EXIT") == "1" {
		Echo(MessageSuccess, "Exit Echo Test", Options{Exit: true})
	}
}
