package tests

import (
	"bytes"
	"log"
	"os"
	
	"testing"
)

const DataDir = "../../tests/data"

func CreateDataDir(t *testing.T) {
	if err := os.MkdirAll(DataDir, 0755); err != nil {
		t.Fatalf("Failed to create test data directory: %v", err)
	}
}

func CaptureOutput(f func()) string {
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	_ = w.Close()
	os.Stdout = old
	_, _ = buf.ReadFrom(r)
	return buf.String()
}

func CaptureLogOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

	f()

	return buf.String()
}