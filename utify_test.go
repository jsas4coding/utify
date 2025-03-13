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

func testEcho(t *testing.T, fn func(string, Options), name string) {
	output := captureOutput(func() {
		fn("Default "+name+" Test", Options{})
	})

	if output == "" {
		t.Errorf("Expected default %s output, got empty string", name)
	}

	output = captureOutput(func() {
		fn("Bold "+name+" Test", Options{Bold: true})
	})

	if output == "" {
		t.Errorf("Expected bold %s output, got empty string", name)
	}

	output = captureOutput(func() {
		fn("Italic "+name+" Test", Options{Italic: true})
	})

	if output == "" {
		t.Errorf("Expected italic %s output, got empty string", name)
	}

	output = captureOutput(func() {
		fn("NoColor "+name+" Test", Options{NoColor: true})
	})

	if output == "" {
		t.Errorf("Expected no color %s output, got empty string", name)
	}

	output = captureOutput(func() {
		fn("NoIcon "+name+" Test", Options{NoIcon: true})
	})

	if output == "" {
		t.Errorf("Expected no icon %s output, got empty string", name)
	}

	if os.Getenv("TEST_"+name+"_EXIT") == "1" {
		fn("Exit "+name+" Test", Options{Exit: true})
	}
}

func TestSuccess(t *testing.T) {
	testEcho(t, Success, "Success")
}

func TestError(t *testing.T) {
	testEcho(t, Error, "Error")
}

func TestWarning(t *testing.T) {
	testEcho(t, Warning, "Warning")
}

func TestInfo(t *testing.T) {
	testEcho(t, Info, "Info")
}

func TestDebug(t *testing.T) {
	testEcho(t, Debug, "Debug")
}

func TestCritical(t *testing.T) {
	testEcho(t, Critical, "Critical")
}

func TestDelete(t *testing.T) {
	testEcho(t, Delete, "Delete")
}

func TestUpdate(t *testing.T) {
	testEcho(t, Update, "Update")
}

func TestInstall(t *testing.T) {
	testEcho(t, Install, "Install")
}

func TestUpgrade(t *testing.T) {
	testEcho(t, Upgrade, "Upgrade")
}

func TestEdit(t *testing.T) {
	testEcho(t, Edit, "Edit")
}

func TestNew(t *testing.T) {
	testEcho(t, New, "New")
}

func TestDownload(t *testing.T) {
	testEcho(t, Download, "Download")
}

func TestUpload(t *testing.T) {
	testEcho(t, Upload, "Upload")
}

func TestSync(t *testing.T) {
	testEcho(t, Sync, "Sync")
}

func TestSearch(t *testing.T) {
	testEcho(t, Search, "Search")
}
