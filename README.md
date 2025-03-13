# Utify

![Go Version](https://img.shields.io/github/go-mod/go-version/jonatas-sas/utify)
![Stable Version](https://img.shields.io/github/v/release/jonatas-sas/utify)
![License](https://img.shields.io/github/license/jonatas-sas/utify)
![Tests](https://github.com/jonatas-sas/utify/actions/workflows/test.yml/badge.svg)
![Code Coverage](https://img.shields.io/codecov/c/github/jonatas-sas/utify)
![Stars](https://img.shields.io/github/stars/jonatas-sas/utify?style=social)

Utify is a Go library for displaying styled messages in the terminal with support for colors, icons, and advanced formatting.

## Installation

To install `utify`, use the following command:

```sh
go get github.com/jonatas-sas/utify
```

## Usage

Basic example:

```go
package main

import (
	"github.com/jonatas-sas/utify"
)

func main() {
	utify.Success("Operation completed successfully!", utify.Options{})
	utify.Error("An unexpected error occurred.", utify.Options{})
	utify.Warning("This might cause issues.", utify.Options{})
	utify.Info("Useful information.", utify.Options{})
	utify.Debug("Debugging enabled.", utify.Options{})
	utify.Critical("Critical error!", utify.Options{})
}
```

## Styling Options

`utify` allows configuring messages with advanced options:

| Option    | Description                        |
| --------- | ---------------------------------- |
| `Bold`    | Displays text in bold              |
| `Italic`  | Displays text in italics           |
| `NoColor` | Removes colors from the output     |
| `NoIcon`  | Removes icons from messages        |
| `Exit`    | Terminates execution after display |

Example using options:

```go
utify.Success("Bold message", utify.Options{Bold: true})
utify.Error("Message without color", utify.Options{NoColor: true})
```

## Available Methods

`utify` provides the following methods to display styled messages:

### **General Status Messages**

- `Success(text string, opts Options)` → Success message
- `Error(text string, opts Options)` → Error message
- `Warning(text string, opts Options)` → Warning message
- `Info(text string, opts Options)` → Informational message
- `Debug(text string, opts Options)` → Debugging message
- `Critical(text string, opts Options)` → Critical error

### **Common Actions**

- `Delete(text string, opts Options)` → Item deletion
- `Update(text string, opts Options)` → Data update
- `Install(text string, opts Options)` → Package installation
- `Upgrade(text string, opts Options)` → Version upgrade
- `Edit(text string, opts Options)` → Item modification
- `New(text string, opts Options)` → Creation of new items

### **Specific Operations**

- `Download(text string, opts Options)` → Indicates a download process
- `Upload(text string, opts Options)` → Indicates an upload process
- `Sync(text string, opts Options)` → Indicates data synchronization
- `Search(text string, opts Options)` → Indicates a search operation

## Using `Echo`

All the functions above internally use the `Echo` method. If you need to customize the output, you can call it directly:

```go
utify.Echo(utify.MessageSuccess, "Custom message", utify.Options{Bold: true, NoIcon: true})
```

## Running Tests

To run the tests, use the following command:

```sh
go test -v
```

The tests verify message output and options to ensure formatting works correctly.

## Contribution

Contributions are welcome! To suggest improvements, open an issue or submit a pull request.

## License

This project is licensed under the MIT license.
