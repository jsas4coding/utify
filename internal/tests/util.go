package tests

import (
	"bytes"
	"log"
	"os"
)

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
