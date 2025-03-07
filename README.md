# Utify

![Build Status](https://github.com/jonatas-sas/utify/actions/workflows/ci_tests.yml/badge.svg)
![Release](https://github.com/jonatas-sas/utify/actions/workflows/ci_release.yml/badge.svg)
![Coverage](https://codecov.io/gh/jonatas-sas/utify/branch/main/graph/badge.svg)

Utify is a Go library for displaying styled and colorized messages in the terminal, with support for icons and flexible configurations.

## 🚀 Installation

To install Utify, run:

```sh
go get github.com/jonatas-sas/utify
```

## ✨ Basic Usage

```go
package main

import (
    "github.com/jonatas-sas/utify"
)

func main() {
    utify.Success("Operation completed successfully", utify.Options{})
    utify.Warning("This is a warning", utify.Options{Bold: true})
}
```

## ⚙️ Available Options

The `Options` struct allows customization of the message appearance:

| Option    | Type   | Description                                         |
| --------- | ------ | --------------------------------------------------- |
| `Bold`    | `bool` | Displays the message in bold.                       |
| `Italic`  | `bool` | Displays the message in italic.                     |
| `NoColor` | `bool` | Disables color output.                              |
| `NoIcon`  | `bool` | Removes the associated icon from the message.       |
| `Exit`    | `bool` | Exits the application after displaying the message. |

## 🖌 Message Types

Utify provides predefined message types, each with an associated color and icon:

```go
utify.Success("Success message", utify.Options{})
utify.Error("Error message", utify.Options{})
utify.Warning("Warning message", utify.Options{})
utify.Info("Info message", utify.Options{})
utify.Debug("Debug message", utify.Options{})
utify.Search("Search message", utify.Options{})
utify.Sync("Sync message", utify.Options{})
utify.Download("Download message", utify.Options{})
utify.Refresh("Refresh message", utify.Options{})
utify.Upload("Upload message", utify.Options{})
utify.Delete("Delete message", utify.Options{})
utify.Critical("Critical message", utify.Options{})
utify.Git("Git message", utify.Options{})
utify.New("New message", utify.Options{})
utify.Edit("Edit message", utify.Options{})
utify.Update("Update message", utify.Options{})
utify.Generation("Generation message", utify.Options{})
utify.Find("Find message", utify.Options{})
utify.Link("Link message", utify.Options{})
utify.Unlink("Unlink message", utify.Options{})
utify.Upgrade("Upgrade message", utify.Options{})
utify.Install("Install message", utify.Options{})
utify.Font("Font message", utify.Options{})
utify.Theme("Theme message", utify.Options{})
utify.Icon("Icon message", utify.Options{})
```

## 🧪 Running Tests

To execute the test suite:

```sh
make test
```

To check test coverage:

```sh
make coverage
```

## 🔍 Linting

To run the linter:

```sh
make lint
```

## 🛠 Customizing Colors

Users can override default colors using `SetColorTable`:

```go
utify.SetColorTable(map[string]string{
    "success": "\033[35m", // Magenta for success messages
})
```

## 📌 Contributing

1. Fork the repository.
2. Create a branch for your feature/fix: `git checkout -b my-feature`
3. Commit your changes: `git commit -m 'My new feature'`
4. Push your branch: `git push origin my-feature`
5. Open a Pull Request 🚀

## 📜 License

This project is licensed under the MIT License. See the `LICENSE` file for details.
