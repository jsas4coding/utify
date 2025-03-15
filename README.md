# Utify

![Go Version](https://img.shields.io/github/go-mod/go-version/jonatas-sas/utify)
![Stable Version](https://img.shields.io/github/v/release/jonatas-sas/utify)
![License](https://img.shields.io/github/license/jonatas-sas/utify)
![Tests](https://github.com/jonatas-sas/utify/actions/workflows/ci.yml/badge.svg)
![Code Coverage](https://img.shields.io/codecov/c/github/jonatas-sas/utify)
![Stars](https://img.shields.io/github/stars/jonatas-sas/utify?style=social)

**Utify** is a Go library for displaying styled messages in the terminal, with support for colors, advanced formatting, and customizable callbacks.

---

## 📌 **Installation**

To install **Utify**, run:

```sh
go get github.com/jonatas-sas/utify
```

---

## 📌 **Basic Usage**

```go
package main

import (
	"github.com/jonatas-sas/utify"
)

func main() {
	opts := utify.OptionsDefault() // Default options

	utify.Success("Operation completed successfully!", opts)
	utify.Error("An unexpected error occurred.", opts)
	utify.Warning("This might cause issues.", opts)
	utify.Info("Useful information.", opts)
	utify.Debug("Debugging enabled.", opts)
	utify.Critical("Critical error!", opts)
}
```

---

## 📌 **Styling Options**

Utify allows configuring messages with various styling options:

| Option     | Description                                                             |
| ---------- | ----------------------------------------------------------------------- |
| `Bold`     | Displays text in **bold**                                               |
| `Italic`   | Displays text in _italic_                                               |
| `NoColor`  | Removes colors from the output                                          |
| `NoIcon`   | Removes icons from messages                                             |
| `NoStyle`  | Removes all formatting (bold/italic)                                    |
| `Exit`     | Terminates execution after displaying the message (disables `Callback`) |
| `Callback` | Executes a function after displaying the message (disables `Exit`)      |

### **Example with options:**

```go
opts := utify.OptionsDefault().
    WithBold().
    WithoutColor()

utify.Success("Bold message without color", opts)
```

---

## 📌 **Using Callbacks**

The `Callback` option allows executing a function after displaying a message. **If `Callback` is set, `Exit` is automatically disabled.** Likewise, **if `Exit` is enabled, `Callback` will be ignored.**

```go
callback := func(msgType utify.MessageType, msg string) {
	fmt.Printf("[Callback] Message displayed: %s - Type: %s\n", msg, msgType)
}

opts := utify.OptionsDefault().
    WithCallback(callback)

utify.Error("Failed to connect to the database!", opts)
```

Output:

```
[31m Failed to connect to the database![0m
[Callback] Message displayed: Failed to connect to the database! - Type: error
```

---

## 📌 **Available Methods**

Utify provides specific methods for different message types:

### **🟢 General Status Messages**

- `Success(text string, opts *Options)` → Success message
- `Error(text string, opts *Options)` → Error message
- `Warning(text string, opts *Options)` → Warning message
- `Info(text string, opts *Options)` → Informational message
- `Debug(text string, opts *Options)` → Debugging message
- `Critical(text string, opts *Options)` → Critical error

### **🛠️ Common Actions**

- `Delete(text string, opts *Options)` → Indicates item deletion
- `Update(text string, opts *Options)` → Indicates data update
- `Install(text string, opts *Options)` → Indicates package installation
- `Upgrade(text string, opts *Options)` → Indicates version upgrade
- `Edit(text string, opts *Options)` → Indicates item modification
- `New(text string, opts *Options)` → Indicates creation of new items

### **📂 Specific Operations**

- `Download(text string, opts *Options)` → Indicates a download process
- `Upload(text string, opts *Options)` → Indicates an upload process
- `Sync(text string, opts *Options)` → Indicates data synchronization
- `Search(text string, opts *Options)` → Indicates a search operation

---

## 📌 **Using `Echo` for Customization**

If you need even more customization, you can call `Echo` directly:

```go
opts := utify.OptionsDefault().
    WithBold().
    WithoutIcon()

utify.Echo(utify.MessageSuccess, "Custom message", opts)
```

---

## 📌 **Running Tests**

To run the tests, use:

```sh
go test -v
```

Tests ensure the correct formatting and behavior of messages and options.

---

## 📌 **Contributing**

Contributions are welcome! To suggest improvements, open an **issue** or submit a **pull request**.

---

## 📌 **License**

This project is licensed under the **MIT License**.
