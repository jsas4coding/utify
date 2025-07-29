# Utify

![Go Version](https://img.shields.io/github/go-mod/go-version/jsas4coding/utify)
![Stable Version](https://img.shields.io/github/v/release/jsas4coding/utify)
![License](https://img.shields.io/github/license/jsas4coding/utify)
![CI/CD](https://img.shields.io/github/actions/workflow/status/jsas4coding/utify/utify.yml?label=CI%2FCD&logo=githubactions&logoColor=white)
![CodeQL](https://img.shields.io/github/actions/workflow/status/jsas4coding/utify/codeql-analysis.yml?label=CodeQL&style=flat-square)
![Coverage](https://img.shields.io/codecov/c/github/jsas4coding/utify?logo=codecov&logoColor=white)
![Stars](https://img.shields.io/github/stars/jsas4coding/utify?style=social)

**Utify** is a powerful Go library for displaying styled messages in the terminal with comprehensive logging support. Features include colored terminal output, structured JSON logging, configurable log targets, and extensive customization options.

---

## 📦 Installation

```sh
go get github.com/jsas4coding/utify
```

---

## 🚀 Basic Usage

```go
package main

import (
	"github.com/jsas4coding/utify"
)

func main() {
	opts := utify.OptionsDefault()

	// Print to terminal AND log to file
	utify.Success("Operation completed!", opts)
	utify.Error("An error occurred!", opts)
	utify.Warning("Pay attention!", opts)
	utify.Info("Here is some info.", opts)
	utify.Debug("Debug message.", opts)
	utify.Critical("System failure!", opts)
}
```

All methods print the message to stdout AND log it to a structured JSON log file. Methods that represent errors (`Error`, `Critical`, `Debug`) return the sentinel `utify.ErrSilent`.

To get the output and handle it manually, use the `Get*` functions:

```go
text, err := utify.GetError("Oops!", opts)
```

---

## 🎨 Styling Options

Customize output using chained methods:

| Method              | Effect                                               |
| ------------------- | ---------------------------------------------------- |
| `.WithBold()`       | Makes the message **bold**                           |
| `.WithItalic()`     | Makes the message _italic_                           |
| `.WithoutColor()`   | Disables all ANSI color codes                        |
| `.WithIcon()`       | Enables icons for messages                           |
| `.WithoutIcon()`    | Disables icons for messages                          |
| `.WithoutStyle()`   | Disables all styling (bold, italic, etc.)            |
| `.WithExit()`       | Exits the program (`os.Exit(1)`) after showing error |
| `.WithCallback(fn)` | Executes callback after message (disables exit)      |

### Example:

```go
opts := utify.OptionsDefault().
  WithBold().
  WithIcon().
  WithoutColor()

utify.Warning("This is bold with icon but no color", opts)
```

---

## 🎯 Icon System

Utify includes a smart icon system with automatic Nerd Font detection and Unicode fallback:

### Features

- **Automatic Detection**: Detects Nerd Font capability in supported terminals
- **Fallback Support**: Uses regular Unicode emoji when Nerd Fonts aren't available
- **Manual Control**: Force specific icon types or disable icons entirely

### Usage

```go
// Enable icons (uses auto-detection)
opts := utify.OptionsDefault().WithIcon()
utify.Success("Operation completed!", opts)

// Disable icons
opts := utify.OptionsDefault().WithoutIcon()
utify.Success("No icon here", opts)
```

### Manual Icon Control

```go
// Force Nerd Font icons (if available in terminal)
utify.ForceNerdFont()

// Force regular Unicode icons
utify.ForceRegularIcons()

// Disable all icons
utify.DisableIcons()

// Check current settings
fmt.Printf("Icon type: %v\n", utify.GetIconType())
fmt.Printf("Nerd Font detected: %v\n", utify.IsNerdFontDetected())
```

### Environment Variable

Set `NERD_FONT_ENABLED=true` to force Nerd Font usage:

```bash
NERD_FONT_ENABLED=true ./my-app
```

### Icon Types

| Icon Type | Description             | Example                  |
| --------- | ----------------------- | ------------------------ |
| Regular   | Unicode emoji (default) | ✅ ❌ ⚠️ ℹ️              |
| Nerd Font | Font Awesome icons      | Various Nerd Font glyphs |
| None      | No icons displayed      | (text only)              |

**Note**: Nerd Font icons require a compatible Nerd Font installed in your terminal. If icons appear blank, your terminal doesn't have the required font glyphs.

---

## 📝 Structured JSON Logging

Utify automatically logs all messages to a structured JSON log file with configurable targets:

### Default Behavior

- **Default location**: `/var/log/{binary_name}.log`
- **Fallback**: Current directory if `/var/log` is not writable
- **Format**: Structured JSON with timestamp, level, message, type, and binary name

### Log Configuration

```go
// Change log file location
err := utify.SetLogTarget("./my-app.log")
if err != nil {
    fmt.Printf("Failed to set log target: %v\n", err)
}

// Get current log location
fmt.Printf("Logging to: %s\n", utify.GetLogTarget())

// Disable/enable logging
utify.SetLoggingEnabled(false)  // Disable logging
utify.SetLoggingEnabled(true)   // Re-enable logging

// Check if logging is enabled
if utify.IsLoggingEnabled() {
    fmt.Println("Logging is active")
}

// Clean up (close log file)
defer utify.CloseLogger()
```

### Log-Only Functions

Use these functions to log messages WITHOUT printing to stdout:

```go
// Log-only functions (no terminal output)
utify.LogSuccess("Operation completed silently")
utify.LogError("Error logged only")
utify.LogInfo("Background info logged")

// Formatted log-only functions
utify.LogSuccessf("Processed %d items", 42)
utify.LogErrorf("Failed to connect to %s", "database")
```

### JSON Log Format

```json
{
  "timestamp": "2025-07-27T16:30:45Z",
  "level": "SUCCESS",
  "message": "Operation completed",
  "type": "success",
  "binary": "my-app"
}
```

---

## 🧐 Using Callbacks

If you want to hook into messages (e.g. for logging, metrics), use `.WithCallback(...)`:

```go
callback := func(t utify.MessageType, msg string) {
	fmt.Printf("\ud83d\udce3 Callback triggered: [%s] %s\n", t, msg)
}

opts := utify.OptionsDefault().
  WithCallback(callback)

utify.Critical("Oops!", opts)
```

️ When `.WithCallback()` is used, `.WithExit()` is ignored — and vice-versa.

---

## 🧹 Available Methods

### 📺 Output Functions (Print + Log)

**General Status**

- `Success(text, opts)`, `Error(text, opts)`, `Warning(text, opts)`
- `Info(text, opts)`, `Debug(text, opts)`, `Critical(text, opts)`

**Common Actions**

- `Delete(text, opts)`, `Update(text, opts)`, `Install(text, opts)`
- `Upgrade(text, opts)`, `Edit(text, opts)`, `New(text, opts)`

**I/O Operations**

- `Download(text, opts)`, `Upload(text, opts)`, `Sync(text, opts)`, `Search(text, opts)`

### 📝 Log-Only Functions (Log Only, No Terminal Output)

**General Status**

- `LogSuccess(text)`, `LogError(text)`, `LogWarning(text)`
- `LogInfo(text)`, `LogDebug(text)`, `LogCritical(text)`

**Common Actions**

- `LogDelete(text)`, `LogUpdate(text)`, `LogInstall(text)`
- `LogUpgrade(text)`, `LogEdit(text)`, `LogNew(text)`

**I/O Operations**

- `LogDownload(text)`, `LogUpload(text)`, `LogSync(text)`, `LogSearch(text)`

### 🔧 Get Functions (Return Values)

- `GetSuccess(text, opts) (string, error)`
- `GetError(text, opts) (string, error)`
- And all other message types...

### 📝 Formatted Versions

Each function has a formatted version with `f` suffix:

```go
// Output + Log
utify.Successf("Success %d: %s", opts, 200, "OK")

// Log only
utify.LogSuccessf("Processed %d items", 42)

// Get formatted
result, err := utify.GetCriticalf("Crash code %d", opts, 42)
```

---

## 🔠 Using Echo (Low-level)

If you need full control, use `Echo(...)`:

```go
opts := utify.OptionsDefault().
  WithBold().
  WithoutIcon()

utify.Echo(utify.MessageInstall, "Installing package...", opts)
```

Return:

```go
(text string, err error)
```

If it's an error-type message, the error returned will be:

```go
utify.ErrSilent
```

---

## 🧪 Testing

**Run all tests:**

```bash
make test
```

**Run specific test suites:**

```bash
make test-unit          # Unit tests only
make test-integration   # Integration tests only
make bench              # Performance benchmarks
```

**Coverage reports:**

```bash
make coverage           # Text coverage report
make coverage-html      # HTML coverage report
```

**Other commands:**

```bash
make build              # Build the project
make lint               # Run linters (requires golangci-lint)
make docs               # Validate documentation
make clean              # Clean up artifacts
```

---

## 📄 License

Licensed under the [MIT License](LICENSE).

---

## 🏗️ Project Structure

```
utify/
├── pkg/                    # Core packages
│   ├── colors/            # ANSI color constants
│   ├── messages/          # Message type definitions
│   ├── options/           # Configuration options
│   ├── formatter/         # Output formatting logic
│   └── logger/            # Structured JSON logging
├── internal/tests/        # Test utilities
├── examples/              # Usage examples
│   ├── basic/            # Basic usage
│   ├── colors/           # Custom colors
│   ├── callbacks/        # Callback functionality
│   └── logging-demo/     # Logging examples
└── tests/                 # Test suites
    ├── unit/             # Unit tests
    ├── integration/      # Integration tests
    └── benchmarks/       # Performance tests
```

## 🚀 What's New in v1.4.0

- **Modular Architecture**: Complete restructure from single-file to organized packages
- **Structured JSON Logging**: Automatic logging to configurable file targets
- **Log-Only Functions**: New `Log*()` functions for logging without terminal output
- **Icon System**: Smart icon support with Nerd Font detection and Unicode fallback
- **Better Testing**: Comprehensive unit, integration, and benchmark tests
- **Enhanced Examples**: Organized examples with dedicated demos
- **Improved Documentation**: Updated with new features and architecture

## 🤝 Contributing

Feel free to open issues, discuss features, or submit PRs! Let's make terminals beautiful, together.

**Development:**

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Run `make test` and `make docs` to ensure quality
5. Submit a pull request
