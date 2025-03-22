# Utify

![Go Version](https://img.shields.io/github/go-mod/go-version/jonatas-sas/utify)
![Stable Version](https://img.shields.io/github/v/release/jonatas-sas/utify)
![License](https://img.shields.io/github/license/jonatas-sas/utify)
![Tests](https://github.com/jonatas-sas/utify/actions/workflows/ci.yml/badge.svg)
![Code Coverage](https://img.shields.io/codecov/c/github/jonatas-sas/utify)
![Stars](https://img.shields.io/github/stars/jonatas-sas/utify?style=social)

**Utify** is a Go library for displaying styled messages in the terminal, with support for color, formatting, error handling, and custom callbacks.

---

## 📦 Installation

```sh
go get github.com/jonatas-sas/utify
```

---

## 🚀 Basic Usage

```go
package main

import (
	"github.com/jonatas-sas/utify"
)

func main() {
	opts := utify.OptionsDefault()

	utify.Success("Operation completed!", opts)
	utify.Error("An error occurred!", opts)
	utify.Warning("Pay attention!", opts)
	utify.Info("Here is some info.", opts)
	utify.Debug("Debug message.", opts)
	utify.Critical("System failure!", opts)
}
```

All methods return `(string, error)`. Errors of type `Error`, `Critical` and `Debug` will return `utify.ErrSilent`, since the message is already displayed on screen.

---

## 🎨 Styling Options

Customize output using chained methods:

| Method              | Effect                                               |
| ------------------- | ---------------------------------------------------- |
| `.WithBold()`       | Makes the message **bold**                           |
| `.WithItalic()`     | Makes the message _italic_                           |
| `.WithoutColor()`   | Disables all ANSI color codes                        |
| `.WithoutIcon()`    | Disables icons (future-proof option)                 |
| `.WithoutStyle()`   | Disables all styling (bold, italic, etc.)            |
| `.WithExit()`       | Exits the program (`os.Exit(1)`) after showing error |
| `.WithCallback(fn)` | Executes callback after message (disables exit)      |

### Example:

```go
opts := utify.OptionsDefault().
  WithBold().
  WithoutColor()

utify.Warning("This is bold but no color", opts)
```

---

## 🧠 Using Callbacks

If you want to hook into messages (e.g. for logging, metrics), use `.WithCallback(...)`.

```go
callback := func(t utify.MessageType, msg string) {
	fmt.Printf("📣 Callback triggered: [%s] %s\n", t, msg)
}

opts := utify.OptionsDefault().
  WithCallback(callback)

utify.Critical("Oops!", opts)
```

⚠️ When `.WithCallback()` is used, `.WithExit()` is ignored — and vice-versa.

---

## 🧹 Available Methods

### ✅ General Status

- `Success(text string, opts *Options)`
- `Error(text string, opts *Options)`
- `Warning(text string, opts *Options)`
- `Info(text string, opts *Options)`
- `Debug(text string, opts *Options)`
- `Critical(text string, opts *Options)`

### 🛠️ Common Actions

- `Delete(...)`, `Update(...)`, `Install(...)`, `Upgrade(...)`, `Edit(...)`, `New(...)`

### 🔄 I/O Operations

- `Download(...)`, `Upload(...)`, `Sync(...)`, `Search(...)`

Each of the above also has a `*f` version, e.g.:

```go
utify.Successf("Success %d: %s", opts, 200, "OK")
```

---

## 💠 Using Echo (Low-level)

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

## ✅ Testing

```bash
go test -v
```

To check coverage:

```bash
go test -coverprofile=cover.out && go tool cover -html=cover.out
```

---

## 📄 License

Licensed under the [MIT License](LICENSE).

---

## 🤝 Contributing

Feel free to open issues, discuss features, or submit PRs! Let's make terminals beautiful, together.
