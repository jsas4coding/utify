# Changelog

## [1.4.3] - 2025-07-29

### 🚀 Features & Improvements

- **CI/CD Pipeline Enhancements**:
  - Unified workflows into `cicd.yml` combining CI, quality checks, examples build, and release automation.
  - Added build process for `examples/` with automatic artifact upload for quick testing.
  - Integrated Codecov coverage reports for main branch pushes.
  - Separated unit, integration, and benchmark tests with clear visibility.
  - Automated Go module publishing to `proxy.golang.org` after tagged releases.

- **Release Automation**:
  - Added automatic changelog generation using `git-chglog`.
  - GitHub Releases now include compiled example binaries.

### 🧪 Quality & Security

- **Enhanced Quality Checks**: Strengthened static analysis with `golangci-lint` and cyclomatic complexity validation.
- **Security**: Removed outdated dependencies in workflows and improved token usage for uploads.

---

## [1.4.2] - 2025-07-27

### ⚡ Performance & Code Quality

- **Code Refactoring**: Reduced cyclomatic complexity for better maintainability
  - Refactored `detectNerdFont()` function: Split complex logic into 4 focused functions
  - Refactored `Echo()` function: Split formatting logic into 6 specialized functions  
  - All functions now have complexity ≤ 9 (under threshold of 10)

### 🔧 CI/CD Improvements

- **Enhanced Workflow Configuration**: Improved GitHub Actions setup
  - **Go Version Management**: All workflows now read Go version from `go.mod` file
  - **Smart CI Triggering**: Optimized to avoid duplicate runs while ensuring comprehensive coverage
    - Runs on: pull requests to main, pushes to main, and tag pushes (releases)
    - Codecov uploads on main branch updates and releases
  - **Updated Codecov Integration**: Upgraded to `codecov/codecov-action@v5` with better error handling
  - **Removed Problematic Dependencies**: Fixed security scanning by removing non-existent packages

### 🏗️ Code Organization

- **Better Separation of Concerns**: Functions now follow single responsibility principle
- **Improved Readability**: Complex logic broken down into smaller, focused functions
- **Enhanced Maintainability**: Changes can be made to specific functionality without affecting others

### 🧪 Quality Assurance

- **Cyclomatic Complexity**: All functions now pass complexity analysis (gocyclo -over 10)
- **Static Analysis**: Clean staticcheck results
- **Comprehensive Testing**: All unit and integration tests continue to pass
- **Lint Compliance**: Maintains 0 lint issues across all linters

---

## [1.4.1] - 2025-07-27

### 🔧 Bug Fixes

- **Lint Compliance**: Fixed all linting issues for clean code quality
  - Fixed 7 `errcheck` violations by properly handling file close operations
  - Fixed 1 `ineffassign` violation by removing unused variable assignment
  - Updated golangci-lint configuration for modern standards

### 🏗️ Code Quality

- **Error Handling**: Improved error handling in logger cleanup operations
- **Test Robustness**: Enhanced test cleanup to properly ignore expected errors
- **Linting**: Achieved 0 lint issues across all linters (errcheck, govet, gosimple, staticcheck, unused, revive, gofmt, goimports)

---

## [1.4.0] - 2025-07-27

### 🚀 Major Features

- **Modular Architecture**: Complete restructure from single-file to proper Go package organization:
  - `pkg/colors/`: ANSI color constants and management
  - `pkg/messages/`: Message types and definitions
  - `pkg/options/`: Options struct with fluent API
  - `pkg/formatter/`: Core formatting logic
  - `pkg/logger/`: Structured JSON logging system
  - `internal/tests/`: Test utilities
  - Organized test structure: unit, integration, and benchmarks

- **Structured JSON Logging**: Full logging system with configurable targets:
  - Default path: `/var/log/{binary_name}.log` (with fallback to current directory)
  - JSON format with timestamp, level, message, type, and binary name
  - Configurable log target: `SetLogTarget(path)`
  - Enable/disable logging: `SetLoggingEnabled(bool)`
  - Log file handle management: `CloseLogger()`

- **Log-Only Functions**: New API for logging without stdout output:
  - `LogSuccess()`, `LogError()`, `LogWarning()`, `LogInfo()`, `LogDebug()`, `LogCritical()`
  - `LogDelete()`, `LogUpdate()`, `LogInstall()`, `LogUpgrade()`, `LogEdit()`, `LogNew()`
  - `LogDownload()`, `LogUpload()`, `LogSync()`, `LogSearch()`
  - Formatted versions: `LogSuccessf()`, `LogErrorf()`, etc.

- **Icon System**: Smart icon display with Nerd Font support:
  - Automatic Nerd Font detection via environment variables (`TERM_PROGRAM`, `NERD_FONT_DETECTED`)
  - Fallback to regular Unicode icons when Nerd Fonts are not available
  - User control: `ForceNerdFont()`, `ForceRegularIcons()`, `DisableIcons()`
  - Integration with existing options: `WithIcon()`, `WithoutIcon()`

### 🔧 Breaking Changes

- **Package Name**: Changed from `github.com/jonatas-sas/utify` to `github.com/jsas4coding/utify`
- **Go Version**: Bumped to Go 1.24.5
- **Test Structure**: Moved `internal/testutil/` to `internal/tests/` with package name change

### ✨ Improvements

- **Fixed Double Message Bug**: Eliminated duplicate message printing
- **Comprehensive Testing**: Added unit tests, integration tests, and benchmarks
- **Better Examples**: Organized examples into separate directories with dedicated demos
- **Enhanced Documentation**: Updated CLAUDE.md with new architecture details
- **Lint Configuration**: Fixed and simplified revive configuration

### 📦 New API Functions

```go
// Logging configuration
SetLogTarget(path string) error
GetLogTarget() string
SetLoggingEnabled(enabled bool)
IsLoggingEnabled() bool
CloseLogger()

// Log-only functions (no stdout output)
LogSuccess(text string)
LogError(text string)
// ... and all other message types

// Formatted log-only functions
LogSuccessf(text string, args ...any)
LogErrorf(text string, args ...any)
// ... and all other message types

// Icon control functions
ForceNerdFont()
ForceRegularIcons()
DisableIcons()
GetIconType() IconType
SetIconType(IconType)
```

### 🧪 Testing

- **Unit Tests**: Complete coverage for all packages
- **Integration Tests**: API compatibility and functionality verification
- **Benchmarks**: Performance testing for all core functions
- **Test Utilities**: Centralized test helpers for output capture

### 📁 Project Structure

```
utify/
├── pkg/                    # Core packages
│   ├── colors/            # ANSI colors
│   ├── messages/          # Message types
│   ├── options/           # Configuration
│   ├── formatter/         # Output formatting
│   └── logger/            # JSON logging
├── internal/tests/        # Test utilities
├── examples/              # Usage examples
│   ├── basic/
│   ├── colors/
│   ├── callbacks/
│   └── logging-demo/
└── tests/                 # Test suites
    ├── unit/
    ├── integration/
    └── benchmarks/
```

---

## [1.3.1] - 2025-03-23

### Added

- **Full GoDoc coverage**: All exported constants, types, functions, and methods now include proper Go-style documentation comments.
- Added `make docs` command to validate documentation using `revive` and `go vet`.
- Improved test readability and structure.

### Changed

- **CI configuration refined**:

  - Now runs only on `pull_request`, `pull_request_target`, and `push` to tags.
  - Removed unnecessary runs on `push` to `main`.
  - Still tests across `ubuntu-latest`, `macos-latest`, and `windows-latest`.

- **Linting setup updated**:

  - Deprecated linters (`deadcode`, `varcheck`, `structcheck`) removed.
  - Migrated from `output.format` → `output.formats` in `golangci-lint` config.
  - Replaced `run.skip-dirs` with `issues.exclude-dirs`.

- Makefile updated:
  - Added `docs` target for lint and vet.
  - Improved log outputs for each target (build, test, clean, etc.).

## [1.3.0] - 2025-03-23

### Added

- Implemented full support for `Get*` and `Get*f` functions that return `(string, error)` for all message types.
- Added tests for `Get*` and `Get*f` variants to ensure correct error handling and output.
- Added `TestSetColorTableOverride` to verify custom ANSI color overrides.
- Added `TestLogOutput` to validate structured logging with message type tags (e.g., `[ERROR]`, `[INFO]`).

### Changed

- Split message function tests to differentiate between return (`Get*`) and no-return (`Success`, `Error`, etc.) implementations.
- Refactored test definitions for proper function signatures and removed invalid error assertions from no-return message methods.
- Improved test clarity and naming consistency across all test cases.

## [1.2.0] - 2025-03-22

### Added

- All message methods (`Success`, `Error`, `Info`, etc.) now return `(string, error)`.
- Introduced `ErrSilent` as a sentinel error for already displayed errors.
- Added fluent option methods for `Options` (e.g., `WithBold`, `WithoutColor`).
- Extended `Echo` to support `Callback` execution or `Exit`, depending on configuration.
- Full test coverage for all message types and formatting combinations.
- Added formatted versions of all message functions (e.g., `Successf`, `Errorf`, etc.).

### Changed

- `Echo` now accepts `*Options` instead of `Options` value.
- `Callback` and `Exit` are now mutually exclusive.
- Improved ANSI handling and color application.

### Fixed

- Fixed handling of `os.Stdout` and `ReadFrom` in test utilities.
- Fixed styling override conflicts between `NoStyle`, `Bold`, and `Italic`.

---

## [1.1.0] - 2025-03-15

### Added

- Introduced `OptionsDefault()` for easier initialization of options.
- Added support for `WithExit()` and `WithCallback()` ensuring mutual exclusivity:
  - If `Callback` is set, `Exit` is automatically disabled.
  - If `Exit` is enabled, `Callback` is removed.
- Implemented **table-driven tests** covering all functionalities.
- Added **log output** for each message type (`Success`, `Error`, `Warning`, etc.).

### Changed

- Improved `Echo()` function to respect `NoColor`, `NoStyle`, and `NoIcon` options correctly.
- Refactored `WithExit()` and `WithCallback()` to properly enforce behavior.
- Enhanced documentation and README to reflect new changes.

### Fixed

- Fixed issue where `Exit` and `Callback` could be set simultaneously.
- Resolved linter error (`errcheck`) by handling the return value of `buf.ReadFrom()` in tests.
- Improved test reliability and robustness.

---

## [1.0.0] - Initial Release

- Basic message styling (`Bold`, `Italic`, `NoColor`, `NoIcon`).
- Predefined message types (`Success`, `Error`, `Warning`, etc.).
- Initial unit tests with basic coverage.
- Custom color mappings using `SetColorTable()`.
