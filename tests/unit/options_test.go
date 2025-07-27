package unit

import (
	"testing"

	"github.com/jsas4coding/utify/pkg/messages"
	"github.com/jsas4coding/utify/pkg/options"
)

func TestOptionsDefault(t *testing.T) {
	opts := options.Default()

	if opts.Bold || opts.Italic || opts.NoColor || opts.NoIcon || opts.NoStyle || opts.Exit {
		t.Error("Default options should have all flags set to false")
	}

	if opts.Callback != nil {
		t.Error("Default options should have nil callback")
	}
}

func TestWithBold(t *testing.T) {
	opts := options.Default().WithBold()

	if !opts.Bold {
		t.Error("WithBold should set Bold to true")
	}
}

func TestWithItalic(t *testing.T) {
	opts := options.Default().WithItalic()

	if !opts.Italic {
		t.Error("WithItalic should set Italic to true")
	}
}

func TestWithoutColor(t *testing.T) {
	opts := options.Default().WithoutColor()

	if !opts.NoColor {
		t.Error("WithoutColor should set NoColor to true")
	}
}

func TestWithoutStyle(t *testing.T) {
	opts := options.Default().WithoutStyle()

	if !opts.NoStyle {
		t.Error("WithoutStyle should set NoStyle to true")
	}
}

func TestWithExit(t *testing.T) {
	opts := options.Default().WithExit()

	if !opts.Exit {
		t.Error("WithExit should set Exit to true")
	}

	if opts.Callback != nil {
		t.Error("WithExit should set Callback to nil")
	}
}

func TestWithCallback(t *testing.T) {
	var called bool
	callback := func(_ messages.Type, _ string) {
		called = true
	}

	opts := options.Default().WithCallback(callback)

	if opts.Callback == nil {
		t.Error("WithCallback should set callback function")
	}

	if opts.Exit {
		t.Error("WithCallback should set Exit to false")
	}

	// Test callback works
	opts.Callback(messages.Success, "test")
	if !called {
		t.Error("Callback should have been called")
	}
}

func TestExitDisablesCallback(t *testing.T) {
	callback := func(_ messages.Type, _ string) {}

	opts := options.Default().WithCallback(callback).WithExit()

	if opts.Callback != nil {
		t.Error("WithExit should disable callback")
	}

	if !opts.Exit {
		t.Error("WithExit should enable exit")
	}
}

func TestCallbackDisablesExit(t *testing.T) {
	callback := func(_ messages.Type, _ string) {}

	opts := options.Default().WithExit().WithCallback(callback)

	if opts.Exit {
		t.Error("WithCallback should disable exit")
	}

	if opts.Callback == nil {
		t.Error("WithCallback should enable callback")
	}
}
