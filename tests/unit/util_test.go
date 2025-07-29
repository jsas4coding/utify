package unit

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/jsas4coding/utify/internal/tests"
)

func TestCaptureOutput(t *testing.T) {
	expected := "Hello, stdout!"
	output := tests.CaptureOutput(func() {
		fmt.Print(expected)
	})
	if !strings.Contains(output, expected) {
		t.Errorf("expected output to contain %q, got %q", expected, output)
	}
}

func TestCaptureLogOutput(t *testing.T) {
	expected := "Hello, log!"
	output := tests.CaptureLogOutput(func() {
		log.Print(expected)
	})
	if !strings.Contains(output, expected) {
		t.Errorf("expected log output to contain %q, got %q", expected, output)
	}
}
