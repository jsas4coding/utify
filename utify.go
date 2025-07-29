/*
Package utify provides a simple and extensible way to display styled messages
in terminal applications, with support for structured logging (JSON),
colored output, and customizable icons.

It is designed to make CLI and TUI applications more readable by providing
predefined message types (success, error, warning, info, etc.), configurable
color tables, and icon sets (including Nerd Font support).

# Features

  - Predefined message types (success, error, warning, info, debug, and more)
  - Direct output functions for quick printing
  - "Get" functions that return formatted strings for further processing
  - Logging-only functions for structured logging without console output
  - Support for Nerd Font and regular icon sets
  - Fully customizable color table
  - Enable/disable structured logging dynamically

# Usage

Basic usage:

	opts := utify.OptionsDefault()
	utify.Success("Operation completed successfully!", opts)
	utify.Error("Something went wrong.", opts)

Formatted output:

	utify.Infof("Hello %s!", opts, "World")

Returning messages without printing:

	msg, err := utify.GetSuccess("Hello!", opts)
	if err == nil {
	    fmt.Println("Formatted message:", msg)
	}

Logging-only (no stdout output):

	utify.SetLoggingEnabled(true)
	utify.LogInfo("This will only be logged, not printed.")

# Customization

Colors and icons can be fully customized:

	utify.SetColorTable(map[string]string{
	    "Red": "#FF0000",
	    "Green": "#00FF00",
	})
	utify.ForceNerdFont()   // Force Nerd Font icons
	utify.DisableIcons()    // Disable icons completely

# Logging

Utify can log messages to a configurable target:

	_ = utify.SetLogTarget("/var/log/myapp.log")
	utify.LogError("An error occurred")

Logging can be enabled or disabled at runtime using:

	utify.SetLoggingEnabled(true)
*/
package utify

import (
	"fmt"

	"github.com/jsas4coding/utify/pkg/colors"
	"github.com/jsas4coding/utify/pkg/formatter"
	"github.com/jsas4coding/utify/pkg/icons"
	"github.com/jsas4coding/utify/pkg/logger"
	"github.com/jsas4coding/utify/pkg/messages"
	"github.com/jsas4coding/utify/pkg/options"
)

// MessageType is an alias for messages.Type for backward compatibility.
type MessageType = messages.Type

// Options is an alias for options.Options for backward compatibility.
type Options = options.Options

var (
	// ErrSilent is returned when output is intentionally silenced.
	ErrSilent = formatter.ErrSilent
	// Echo formats and prints a message to the terminal.
	Echo = formatter.Echo

	// Predefined message types for quick use.
	MessageSuccess    = messages.Success
	MessageError      = messages.Error
	MessageWarning    = messages.Warning
	MessageInfo       = messages.Info
	MessageDebug      = messages.Debug
	MessageSearch     = messages.Search
	MessageSync       = messages.Sync
	MessageDownload   = messages.Download
	MessageRefresh    = messages.Refresh
	MessageUpload     = messages.Upload
	MessageDelete     = messages.Delete
	MessageCritical   = messages.Critical
	MessageGit        = messages.Git
	MessageNew        = messages.New
	MessageEdit       = messages.Edit
	MessageUpdate     = messages.Update
	MessageGeneration = messages.Generation
	MessageFind       = messages.Find
	MessageLink       = messages.Link
	MessageUnlink     = messages.Unlink
	MessageUpgrade    = messages.Upgrade
	MessageInstall    = messages.Install
	MessageFont       = messages.Font
	MessageTheme      = messages.Theme
	MessageIcon       = messages.Icon
	Default           = messages.Default

	// Color and style constants for text formatting.
	ColorRed       = colors.Red
	ColorGreen     = colors.Green
	ColorYellow    = colors.Yellow
	ColorBlue      = colors.Blue
	ColorCyan      = colors.Cyan
	ColorWhite     = colors.White
	ColorMagenta   = colors.Magenta
	ColorGray      = colors.Gray
	ColorLightBlue = colors.LightBlue
	StyleBold      = colors.Bold
	StyleItalic    = colors.Italic
	StyleReset     = colors.Reset
)

// OptionsDefault returns the default configuration for message output.
func OptionsDefault() *Options {
	return options.Default()
}

// SetColorTable overrides the default color table with a custom one.
func SetColorTable(newColors map[string]string) {
	colors.SetColorTable(newColors)
}

// SetLogTarget sets the destination for structured logs (e.g., file path or stdout).
func SetLogTarget(target string) error {
	return logger.SetLogTarget(target)
}

// GetLogTarget returns the current log destination.
func GetLogTarget() string {
	return logger.GetLogTarget()
}

// SetLoggingEnabled enables or disables structured logging.
func SetLoggingEnabled(enabled bool) {
	logger.SetEnabled(enabled)
}

// IsLoggingEnabled checks whether structured logging is enabled.
func IsLoggingEnabled() bool {
	return logger.IsEnabled()
}

// CloseLogger closes any active log writers.
func CloseLogger() {
	logger.Close()
}

// ForceNerdFont forces icon rendering to use Nerd Fonts.
func ForceNerdFont() {
	icons.ForceNerdFont()
}

// ForceRegularIcons forces icon rendering to use regular (non-Nerd Font) icons.
func ForceRegularIcons() {
	icons.ForceRegularIcons()
}

// DisableIcons disables icon rendering.
func DisableIcons() {
	icons.DisableIcons()
}

// IsNerdFontDetected returns whether a Nerd Font is detected in the environment.
func IsNerdFontDetected() bool {
	return icons.IsNerdFontDetected()
}

// GetIconType returns the current icon type (as an integer enum).
func GetIconType() int {
	return int(icons.GetIconType())
}

// --- Direct output functions (print to stdout) ---

// Success prints a success message to stdout.
func Success(text string, opts *Options) {
	_, _ = Echo(MessageSuccess, text, opts)
}

// Error prints an error message to stdout.
func Error(text string, opts *Options) {
	_, _ = Echo(MessageError, text, opts)
}

// Warning prints a warning message to stdout.
func Warning(text string, opts *Options) {
	_, _ = Echo(MessageWarning, text, opts)
}

// Info prints an info message to stdout.
func Info(text string, opts *Options) {
	_, _ = Echo(MessageInfo, text, opts)
}

// Debug prints a debug message to stdout.
func Debug(text string, opts *Options) {
	_, _ = Echo(MessageDebug, text, opts)
}

// Critical prints a critical message to stdout.
func Critical(text string, opts *Options) {
	_, _ = Echo(MessageCritical, text, opts)
}

// Delete prints a delete message to stdout.
func Delete(text string, opts *Options) {
	_, _ = Echo(MessageDelete, text, opts)
}

// Update prints an update message to stdout.
func Update(text string, opts *Options) {
	_, _ = Echo(MessageUpdate, text, opts)
}

// Install prints an install message to stdout.
func Install(text string, opts *Options) {
	_, _ = Echo(MessageInstall, text, opts)
}

// Upgrade prints an upgrade message to stdout.
func Upgrade(text string, opts *Options) {
	_, _ = Echo(MessageUpgrade, text, opts)
}

// Edit prints an edit message to stdout.
func Edit(text string, opts *Options) {
	_, _ = Echo(MessageEdit, text, opts)
}

// New prints a new message to stdout.
func New(text string, opts *Options) {
	_, _ = Echo(MessageNew, text, opts)
}

// Download prints a download message to stdout.
func Download(text string, opts *Options) {
	_, _ = Echo(MessageDownload, text, opts)
}

// Upload prints an upload message to stdout.
func Upload(text string, opts *Options) {
	_, _ = Echo(MessageUpload, text, opts)
}

// Sync prints a sync message to stdout.
func Sync(text string, opts *Options) {
	_, _ = Echo(MessageSync, text, opts)
}

// Search prints a search message to stdout.
func Search(text string, opts *Options) {
	_, _ = Echo(MessageSearch, text, opts)
}

// --- Formatted direct output functions (printf style) ---

// Successf prints a formatted success message to stdout.
func Successf(text string, opts *Options, args ...any) {
	Success(fmt.Sprintf(text, args...), opts)
}

// Errorf prints a formatted error message to stdout.
func Errorf(text string, opts *Options, args ...any) {
	Error(fmt.Sprintf(text, args...), opts)
}

// Warningf prints a formatted warning message to stdout.
func Warningf(text string, opts *Options, args ...any) {
	Warning(fmt.Sprintf(text, args...), opts)
}

// Infof prints a formatted info message to stdout.
func Infof(text string, opts *Options, args ...any) {
	Info(fmt.Sprintf(text, args...), opts)
}

// Debugf prints a formatted debug message to stdout.
func Debugf(text string, opts *Options, args ...any) {
	Debug(fmt.Sprintf(text, args...), opts)
}

// Criticalf prints a formatted critical message to stdout.
func Criticalf(text string, opts *Options, args ...any) {
	Critical(fmt.Sprintf(text, args...), opts)
}

// Deletef prints a formatted delete message to stdout.
func Deletef(text string, opts *Options, args ...any) {
	Delete(fmt.Sprintf(text, args...), opts)
}

// Updatef prints a formatted update message to stdout.
func Updatef(text string, opts *Options, args ...any) {
	Update(fmt.Sprintf(text, args...), opts)
}

// Installf prints a formatted install message to stdout.
func Installf(text string, opts *Options, args ...any) {
	Install(fmt.Sprintf(text, args...), opts)
}

// Upgradef prints a formatted upgrade message to stdout.
func Upgradef(text string, opts *Options, args ...any) {
	Upgrade(fmt.Sprintf(text, args...), opts)
}

// Editf prints a formatted edit message to stdout.
func Editf(text string, opts *Options, args ...any) {
	Edit(fmt.Sprintf(text, args...), opts)
}

// Newf prints a formatted new message to stdout.
func Newf(text string, opts *Options, args ...any) {
	New(fmt.Sprintf(text, args...), opts)
}

// Downloadf prints a formatted download message to stdout.
func Downloadf(text string, opts *Options, args ...any) {
	Download(fmt.Sprintf(text, args...), opts)
}

// Uploadf prints a formatted upload message to stdout.
func Uploadf(text string, opts *Options, args ...any) {
	Upload(fmt.Sprintf(text, args...), opts)
}

// Syncf prints a formatted sync message to stdout.
func Syncf(text string, opts *Options, args ...any) {
	Sync(fmt.Sprintf(text, args...), opts)
}

// Searchf prints a formatted search message to stdout.
func Searchf(text string, opts *Options, args ...any) {
	Search(fmt.Sprintf(text, args...), opts)
}

// --- Get functions (return formatted strings instead of printing) ---

// GetSuccess returns a formatted success message as a string.
func GetSuccess(text string, opts *Options) (string, error) {
	return Echo(MessageSuccess, text, opts)
}

// GetError returns a formatted error message as a string.
func GetError(text string, opts *Options) (string, error) {
	return Echo(MessageError, text, opts)
}

// GetWarning returns a formatted warning message as a string.
func GetWarning(text string, opts *Options) (string, error) {
	return Echo(MessageWarning, text, opts)
}

// GetInfo returns a formatted info message as a string.
func GetInfo(text string, opts *Options) (string, error) {
	return Echo(MessageInfo, text, opts)
}

// GetDebug returns a formatted debug message as a string.
func GetDebug(text string, opts *Options) (string, error) {
	return Echo(MessageDebug, text, opts)
}

// GetCritical returns a formatted critical message as a string.
func GetCritical(text string, opts *Options) (string, error) {
	return Echo(MessageCritical, text, opts)
}

// GetDelete returns a formatted delete message as a string.
func GetDelete(text string, opts *Options) (string, error) {
	return Echo(MessageDelete, text, opts)
}

// GetUpdate returns a formatted update message as a string.
func GetUpdate(text string, opts *Options) (string, error) {
	return Echo(MessageUpdate, text, opts)
}

// GetInstall returns a formatted install message as a string.
func GetInstall(text string, opts *Options) (string, error) {
	return Echo(MessageInstall, text, opts)
}

// GetUpgrade returns a formatted upgrade message as a string.
func GetUpgrade(text string, opts *Options) (string, error) {
	return Echo(MessageUpgrade, text, opts)
}

// GetEdit returns a formatted edit message as a string.
func GetEdit(text string, opts *Options) (string, error) {
	return Echo(MessageEdit, text, opts)
}

// GetNew returns a formatted new message as a string.
func GetNew(text string, opts *Options) (string, error) {
	return Echo(MessageNew, text, opts)
}

// GetDownload returns a formatted download message as a string.
func GetDownload(text string, opts *Options) (string, error) {
	return Echo(MessageDownload, text, opts)
}

// GetUpload returns a formatted upload message as a string.
func GetUpload(text string, opts *Options) (string, error) {
	return Echo(MessageUpload, text, opts)
}

// GetSync returns a formatted sync message as a string.
func GetSync(text string, opts *Options) (string, error) {
	return Echo(MessageSync, text, opts)
}

// GetSearch returns a formatted search message as a string.
func GetSearch(text string, opts *Options) (string, error) {
	return Echo(MessageSearch, text, opts)
}

// --- Get formatted functions (printf style) ---

// GetSuccessf returns a formatted success message with arguments.
func GetSuccessf(text string, opts *Options, args ...any) (string, error) {
	return GetSuccess(fmt.Sprintf(text, args...), opts)
}

// GetErrorf returns a formatted error message with arguments.
func GetErrorf(text string, opts *Options, args ...any) (string, error) {
	return GetError(fmt.Sprintf(text, args...), opts)
}

// GetWarningf returns a formatted warning message with arguments.
func GetWarningf(text string, opts *Options, args ...any) (string, error) {
	return GetWarning(fmt.Sprintf(text, args...), opts)
}

// GetInfof returns a formatted info message with arguments.
func GetInfof(text string, opts *Options, args ...any) (string, error) {
	return GetInfo(fmt.Sprintf(text, args...), opts)
}

// GetDebugf returns a formatted debug message with arguments.
func GetDebugf(text string, opts *Options, args ...any) (string, error) {
	return GetDebug(fmt.Sprintf(text, args...), opts)
}

// GetCriticalf returns a formatted critical message with arguments.
func GetCriticalf(text string, opts *Options, args ...any) (string, error) {
	return GetCritical(fmt.Sprintf(text, args...), opts)
}

// GetDeletef returns a formatted delete message with arguments.
func GetDeletef(text string, opts *Options, args ...any) (string, error) {
	return GetDelete(fmt.Sprintf(text, args...), opts)
}

// GetUpdatef returns a formatted update message with arguments.
func GetUpdatef(text string, opts *Options, args ...any) (string, error) {
	return GetUpdate(fmt.Sprintf(text, args...), opts)
}

// GetInstallf returns a formatted install message with arguments.
func GetInstallf(text string, opts *Options, args ...any) (string, error) {
	return GetInstall(fmt.Sprintf(text, args...), opts)
}

// GetUpgradef returns a formatted upgrade message with arguments.
func GetUpgradef(text string, opts *Options, args ...any) (string, error) {
	return GetUpgrade(fmt.Sprintf(text, args...), opts)
}

// GetEditf returns a formatted edit message with arguments.
func GetEditf(text string, opts *Options, args ...any) (string, error) {
	return GetEdit(fmt.Sprintf(text, args...), opts)
}

// GetNewf returns a formatted new message with arguments.
func GetNewf(text string, opts *Options, args ...any) (string, error) {
	return GetNew(fmt.Sprintf(text, args...), opts)
}

// GetDownloadf returns a formatted download message with arguments.
func GetDownloadf(text string, opts *Options, args ...any) (string, error) {
	return GetDownload(fmt.Sprintf(text, args...), opts)
}

// GetUploadf returns a formatted upload message with arguments.
func GetUploadf(text string, opts *Options, args ...any) (string, error) {
	return GetUpload(fmt.Sprintf(text, args...), opts)
}

// GetSyncf returns a formatted sync message with arguments.
func GetSyncf(text string, opts *Options, args ...any) (string, error) {
	return GetSync(fmt.Sprintf(text, args...), opts)
}

// GetSearchf returns a formatted search message with arguments.
func GetSearchf(text string, opts *Options, args ...any) (string, error) {
	return GetSearch(fmt.Sprintf(text, args...), opts)
}

// --- Log-only functions (structured logging only, no stdout) ---

// LogSuccess logs a success message without printing to stdout.
func LogSuccess(text string) {
	logger.LogOnly(MessageSuccess, text)
}

// LogError logs an error message without printing to stdout.
func LogError(text string) {
	logger.LogOnly(MessageError, text)
}

// LogWarning logs a warning message without printing to stdout.
func LogWarning(text string) {
	logger.LogOnly(MessageWarning, text)
}

// LogInfo logs an info message without printing to stdout.
func LogInfo(text string) {
	logger.LogOnly(MessageInfo, text)
}

// LogDebug logs a debug message without printing to stdout.
func LogDebug(text string) {
	logger.LogOnly(MessageDebug, text)
}

// LogCritical logs a critical message without printing to stdout.
func LogCritical(text string) {
	logger.LogOnly(MessageCritical, text)
}

// LogDelete logs a delete message without printing to stdout.
func LogDelete(text string) {
	logger.LogOnly(MessageDelete, text)
}

// LogUpdate logs an update message without printing to stdout.
func LogUpdate(text string) {
	logger.LogOnly(MessageUpdate, text)
}

// LogInstall logs an install message without printing to stdout.
func LogInstall(text string) {
	logger.LogOnly(MessageInstall, text)
}

// LogUpgrade logs an upgrade message without printing to stdout.
func LogUpgrade(text string) {
	logger.LogOnly(MessageUpgrade, text)
}

// LogEdit logs an edit message without printing to stdout.
func LogEdit(text string) {
	logger.LogOnly(MessageEdit, text)
}

// LogNew logs a new message without printing to stdout.
func LogNew(text string) {
	logger.LogOnly(MessageNew, text)
}

// LogDownload logs a download message without printing to stdout.
func LogDownload(text string) {
	logger.LogOnly(MessageDownload, text)
}

// LogUpload logs an upload message without printing to stdout.
func LogUpload(text string) {
	logger.LogOnly(MessageUpload, text)
}

// LogSync logs a sync message without printing to stdout.
func LogSync(text string) {
	logger.LogOnly(MessageSync, text)
}

// LogSearch logs a search message without printing to stdout.
func LogSearch(text string) {
	logger.LogOnly(MessageSearch, text)
}

// --- Log-only formatted functions (printf style) ---

// LogSuccessf logs a formatted success message without printing to stdout.
func LogSuccessf(text string, args ...any) {
	LogSuccess(fmt.Sprintf(text, args...))
}

// LogErrorf logs a formatted error message without printing to stdout.
func LogErrorf(text string, args ...any) {
	LogError(fmt.Sprintf(text, args...))
}

// LogWarningf logs a formatted warning message without printing to stdout.
func LogWarningf(text string, args ...any) {
	LogWarning(fmt.Sprintf(text, args...))
}

// LogInfof logs a formatted info message without printing to stdout.
func LogInfof(text string, args ...any) {
	LogInfo(fmt.Sprintf(text, args...))
}

// LogDebugf logs a formatted debug message without printing to stdout.
func LogDebugf(text string, args ...any) {
	LogDebug(fmt.Sprintf(text, args...))
}

// LogCriticalf logs a formatted critical message without printing to stdout.
func LogCriticalf(text string, args ...any) {
	LogCritical(fmt.Sprintf(text, args...))
}

// LogDeletef logs a formatted delete message without printing to stdout.
func LogDeletef(text string, args ...any) {
	LogDelete(fmt.Sprintf(text, args...))
}

// LogUpdatef logs a formatted update message without printing to stdout.
func LogUpdatef(text string, args ...any) {
	LogUpdate(fmt.Sprintf(text, args...))
}

// LogInstallf logs a formatted install message without printing to stdout.
func LogInstallf(text string, args ...any) {
	LogInstall(fmt.Sprintf(text, args...))
}

// LogUpgradef logs a formatted upgrade message without printing to stdout.
func LogUpgradef(text string, args ...any) {
	LogUpgrade(fmt.Sprintf(text, args...))
}

// LogEditf logs a formatted edit message without printing to stdout.
func LogEditf(text string, args ...any) {
	LogEdit(fmt.Sprintf(text, args...))
}

// LogNewf logs a formatted new message without printing to stdout.
func LogNewf(text string, args ...any) {
	LogNew(fmt.Sprintf(text, args...))
}

// LogDownloadf logs a formatted download message without printing to stdout.
func LogDownloadf(text string, args ...any) {
	LogDownload(fmt.Sprintf(text, args...))
}

// LogUploadf logs a formatted upload message without printing to stdout.
func LogUploadf(text string, args ...any) {
	LogUpload(fmt.Sprintf(text, args...))
}

// LogSyncf logs a formatted sync message without printing to stdout.
func LogSyncf(text string, args ...any) {
	LogSync(fmt.Sprintf(text, args...))
}

// LogSearchf logs a formatted search message without printing to stdout.
func LogSearchf(text string, args ...any) {
	LogSearch(fmt.Sprintf(text, args...))
}
