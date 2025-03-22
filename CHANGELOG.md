# Changelog

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
