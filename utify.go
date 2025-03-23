// Package utify provides styled terminal output with color, formatting, and behavior control.
package utify

import (
	"errors"
	"fmt"
	"log"
	"maps"
	"os"
	"strings"
)

// ANSI color codes and styles for terminal output
const (
	ColorRed       = "\033[31m"
	ColorGreen     = "\033[32m"
	ColorYellow    = "\033[33m"
	ColorBlue      = "\033[34m"
	ColorCyan      = "\033[36m"
	ColorWhite     = "\033[97m"
	ColorMagenta   = "\033[35m"
	ColorGray      = "\033[90m"
	ColorLightBlue = "\033[94m"
	StyleBold      = "\033[1m"
	StyleItalic    = "\033[3m"
	StyleReset     = "\033[0m"
)

// MessageType defines supported message categories for terminal output.
type MessageType string

// Options defines configuration flags for how messages are styled and handled.
type Options struct {
	Bold     bool
	Italic   bool
	NoColor  bool
	NoIcon   bool
	NoStyle  bool
	Exit     bool
	Callback func(MessageType, string)
}

// OptionsDefault returns a new Options instance with default configuration.
func OptionsDefault() *Options {
	return &Options{}
}

// WithBold enables bold formatting for the message.
func (o *Options) WithBold() *Options {
	o.Bold = true
	return o
}

// WithItalic enables italic formatting for the message.
func (o *Options) WithItalic() *Options {
	o.Italic = true
	return o
}

// WithoutColor disables all color formatting for the message.
func (o *Options) WithoutColor() *Options {
	o.NoColor = true
	return o
}

// WithoutIcon disables the icon (future-proof option).
func (o *Options) WithoutIcon() *Options {
	o.NoIcon = true
	return o
}

// WithoutStyle disables bold and italic formatting.
func (o *Options) WithoutStyle() *Options {
	o.NoStyle = true
	return o
}

// WithExit sets the option to exit the program after printing an error message.
func (o *Options) WithExit() *Options {
	o.Exit = true
	o.Callback = nil
	return o
}

// WithCallback sets a custom callback to run after the message is printed.
// This disables automatic program exit.
func (o *Options) WithCallback(cb func(MessageType, string)) *Options {
	o.Callback = cb
	o.Exit = false
	return o
}

// MessageType constants represent predefined categories for terminal messages.
const (
	MessageSuccess    MessageType = "success"
	MessageError      MessageType = "error"
	MessageWarning    MessageType = "warning"
	MessageInfo       MessageType = "info"
	MessageDebug      MessageType = "debug"
	MessageSearch     MessageType = "search"
	MessageSync       MessageType = "sync"
	MessageDownload   MessageType = "download"
	MessageRefresh    MessageType = "refresh"
	MessageUpload     MessageType = "upload"
	MessageDelete     MessageType = "delete"
	MessageCritical   MessageType = "critical"
	MessageGit        MessageType = "git"
	MessageNew        MessageType = "new"
	MessageEdit       MessageType = "edit"
	MessageUpdate     MessageType = "update"
	MessageGeneration MessageType = "generation"
	MessageFind       MessageType = "find"
	MessageLink       MessageType = "link"
	MessageUnlink     MessageType = "unlink"
	MessageUpgrade    MessageType = "upgrade"
	MessageInstall    MessageType = "install"
	MessageFont       MessageType = "font"
	MessageTheme      MessageType = "theme"
	MessageIcon       MessageType = "icon"
	Default           MessageType = "default"
)

// colors maps each MessageType to its corresponding ANSI color.
var colors = map[MessageType]string{
	MessageSuccess:    ColorGreen,
	MessageError:      ColorRed,
	MessageWarning:    ColorYellow,
	MessageInfo:       ColorCyan,
	MessageDebug:      ColorGray,
	MessageSearch:     ColorBlue,
	MessageSync:       ColorMagenta,
	MessageDownload:   ColorWhite,
	MessageRefresh:    ColorLightBlue,
	MessageUpload:     ColorGreen,
	MessageDelete:     ColorRed,
	MessageCritical:   ColorMagenta,
	MessageGit:        ColorMagenta,
	MessageNew:        ColorGreen,
	MessageEdit:       ColorBlue,
	MessageUpdate:     ColorYellow,
	MessageGeneration: ColorCyan,
	MessageFind:       ColorBlue,
	MessageLink:       ColorMagenta,
	MessageUnlink:     ColorRed,
	MessageUpgrade:    ColorLightBlue,
	MessageInstall:    ColorGreen,
	MessageFont:       ColorWhite,
	MessageTheme:      ColorMagenta,
	MessageIcon:       ColorWhite,
	Default:           ColorWhite,
}

// ErrSilent is returned for errors that have already been displayed.
var ErrSilent = errors.New("silent error")

var userColors = map[string]string{}

// getColor returns the ANSI color for a given message type,
// checking user overrides before falling back to defaults.
func getColor(msgType MessageType) string {
	if color, exists := userColors[string(msgType)]; exists {
		return color
	}

	return colors[msgType]
}

// SetColorTable replaces or extends the internal color map.
func SetColorTable(newColors map[string]string) {
	maps.Copy(userColors, newColors)
}

// Echo renders a styled message of the given type and returns the raw message and an optional error.
// If a callback is set in options, it will be executed. If Exit is true and the message is an error type, the program will exit.
func Echo(msgType MessageType, text string, opts *Options) (string, error) {
	color := getColor(msgType)

	if opts.NoColor {
		color = ""
	}

	style := ""

	if opts.Bold {
		style += StyleBold
	}

	if opts.Italic {
		style += StyleItalic
	}

	if opts.NoStyle {
		style = ""
	}

	message := fmt.Sprintf("%s%s %s%s", style, color, text, StyleReset)

	fmt.Println(message)

	log.Printf("[%s] %s", strings.ToUpper(string(msgType)), text)

	errorTypes := map[MessageType]bool{
		MessageError:    true,
		MessageCritical: true,
		MessageDebug:    true,
	}

	if opts.Callback != nil {
		opts.Callback(msgType, text)
	} else if opts.Exit && errorTypes[msgType] {
		os.Exit(1)
	}

	if errorTypes[msgType] {
		return text, ErrSilent
	}

	return text, nil
}

// Success prints a green success message to the terminal.
func Success(text string, opts *Options) {
	_, _ = Echo(MessageSuccess, text, opts)
}

// Error prints a red error message to the terminal.
func Error(text string, opts *Options) {
	_, _ = Echo(MessageError, text, opts)
}

// Warning prints a yellow warning message to the terminal.
func Warning(text string, opts *Options) {
	_, _ = Echo(MessageWarning, text, opts)
}

// Info prints a cyan informational message to the terminal.
func Info(text string, opts *Options) {
	_, _ = Echo(MessageInfo, text, opts)
}

// Debug prints a gray debug message to the terminal.
func Debug(text string, opts *Options) {
	_, _ = Echo(MessageDebug, text, opts)
}

// Critical prints a magenta critical error message to the terminal.
func Critical(text string, opts *Options) {
	_, _ = Echo(MessageCritical, text, opts)
}

// Delete prints a red delete action message to the terminal.
func Delete(text string, opts *Options) {
	_, _ = Echo(MessageDelete, text, opts)
}

// Update prints a yellow update action message to the terminal.
func Update(text string, opts *Options) {
	_, _ = Echo(MessageUpdate, text, opts)
}

// Install prints a green install action message to the terminal.
func Install(text string, opts *Options) {
	_, _ = Echo(MessageInstall, text, opts)
}

// Upgrade prints a light blue upgrade action message to the terminal.
func Upgrade(text string, opts *Options) {
	_, _ = Echo(MessageUpgrade, text, opts)
}

// Edit prints a blue edit action message to the terminal.
func Edit(text string, opts *Options) {
	_, _ = Echo(MessageEdit, text, opts)
}

// New prints a green message for newly created items.
func New(text string, opts *Options) {
	_, _ = Echo(MessageNew, text, opts)
}

// Download prints a message related to downloads.
func Download(text string, opts *Options) {
	_, _ = Echo(MessageDownload, text, opts)
}

// Upload prints a message related to uploads.
func Upload(text string, opts *Options) {
	_, _ = Echo(MessageUpload, text, opts)
}

// Sync prints a message related to synchronization.
func Sync(text string, opts *Options) {
	_, _ = Echo(MessageSync, text, opts)
}

// Search prints a message related to search operations.
func Search(text string, opts *Options) {
	_, _ = Echo(MessageSearch, text, opts)
}

// Successf formats and prints a success message.
func Successf(text string, opts *Options, args ...any) {
	text = fmt.Sprintf(text, args...)
	Success(text, opts)
}

// Errorf formats and prints an error message.
func Errorf(text string, opts *Options, args ...any) {
	text = fmt.Sprintf(text, args...)
	Error(text, opts)
}

// Warningf formats and prints a warning message.
func Warningf(text string, opts *Options, args ...any) {
	text = fmt.Sprintf(text, args...)
	Warning(text, opts)
}

// Infof formats and prints an informational message.
func Infof(text string, opts *Options, args ...any) {
	text = fmt.Sprintf(text, args...)
	Info(text, opts)
}

// Debugf formats and prints a debug message.
func Debugf(text string, opts *Options, args ...any) {
	text = fmt.Sprintf(text, args...)
	Debug(text, opts)
}

// Criticalf formats and prints a critical error message.
func Criticalf(text string, opts *Options, args ...any) {
	text = fmt.Sprintf(text, args...)
	Critical(text, opts)
}

// Deletef formats and prints a delete message.
func Deletef(text string, opts *Options, args ...any) {
	text = fmt.Sprintf(text, args...)
	Delete(text, opts)
}

// Updatef formats and prints an update message.
func Updatef(text string, opts *Options, args ...any) {
	text = fmt.Sprintf(text, args...)
	Update(text, opts)
}

// Installf formats and prints an install message.
func Installf(text string, opts *Options, args ...any) {
	text = fmt.Sprintf(text, args...)
	Install(text, opts)
}

// Upgradef formats and prints an upgrade message.
func Upgradef(text string, opts *Options, args ...any) {
	text = fmt.Sprintf(text, args...)
	Upgrade(text, opts)
}

// Editf formats and prints an edit message.
func Editf(text string, opts *Options, args ...any) {
	text = fmt.Sprintf(text, args...)
	Edit(text, opts)
}

// Newf formats and prints a creation message.
func Newf(text string, opts *Options, args ...any) {
	text = fmt.Sprintf(text, args...)
	New(text, opts)
}

// Downloadf formats and prints a download message.
func Downloadf(text string, opts *Options, args ...any) {
	text = fmt.Sprintf(text, args...)
	Download(text, opts)
}

// Uploadf formats and prints an upload message.
func Uploadf(text string, opts *Options, args ...any) {
	text = fmt.Sprintf(text, args...)
	Upload(text, opts)
}

// Syncf formats and prints a sync message.
func Syncf(text string, opts *Options, args ...any) {
	text = fmt.Sprintf(text, args...)
	Sync(text, opts)
}

// Searchf formats and prints a search message.
func Searchf(text string, opts *Options, args ...any) {
	text = fmt.Sprintf(text, args...)
	Search(text, opts)
}

// GetSuccess returns a success message with optional error.
func GetSuccess(text string, opts *Options) (string, error) {
	return Echo(MessageSuccess, text, opts)
}

// GetError returns an error message with optional error.
func GetError(text string, opts *Options) (string, error) {
	return Echo(MessageError, text, opts)
}

// GetWarning returns a warning message with optional error.
func GetWarning(text string, opts *Options) (string, error) {
	return Echo(MessageWarning, text, opts)
}

// GetInfo returns an informational message with optional error.
func GetInfo(text string, opts *Options) (string, error) {
	return Echo(MessageInfo, text, opts)
}

// GetDebug returns a debug message with optional error.
func GetDebug(text string, opts *Options) (string, error) {
	return Echo(MessageDebug, text, opts)
}

// GetCritical returns a critical message with optional error.
func GetCritical(text string, opts *Options) (string, error) {
	return Echo(MessageCritical, text, opts)
}

// GetDelete returns a delete message with optional error.
func GetDelete(text string, opts *Options) (string, error) {
	return Echo(MessageDelete, text, opts)
}

// GetUpdate returns an update message with optional error.
func GetUpdate(text string, opts *Options) (string, error) {
	return Echo(MessageUpdate, text, opts)
}

// GetInstall returns an install message with optional error.
func GetInstall(text string, opts *Options) (string, error) {
	return Echo(MessageInstall, text, opts)
}

// GetUpgrade returns an upgrade message with optional error.
func GetUpgrade(text string, opts *Options) (string, error) {
	return Echo(MessageUpgrade, text, opts)
}

// GetEdit returns an edit message with optional error.
func GetEdit(text string, opts *Options) (string, error) {
	return Echo(MessageEdit, text, opts)
}

// GetNew returns a creation message with optional error.
func GetNew(text string, opts *Options) (string, error) {
	return Echo(MessageNew, text, opts)
}

// GetDownload returns a download message with optional error.
func GetDownload(text string, opts *Options) (string, error) {
	return Echo(MessageDownload, text, opts)
}

// GetUpload returns an upload message with optional error.
func GetUpload(text string, opts *Options) (string, error) {
	return Echo(MessageUpload, text, opts)
}

// GetSync returns a sync message with optional error.
func GetSync(text string, opts *Options) (string, error) {
	return Echo(MessageSync, text, opts)
}

// GetSearch returns a search message with optional error.
func GetSearch(text string, opts *Options) (string, error) {
	return Echo(MessageSearch, text, opts)
}

// GetSuccessf returns a formatted success message with optional error.
func GetSuccessf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)
	return GetSuccess(text, opts)
}

// GetErrorf returns a formatted error message with optional error.
func GetErrorf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)
	return GetError(text, opts)
}

// GetWarningf returns a formatted warning message with optional error.
func GetWarningf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)
	return GetWarning(text, opts)
}

// GetInfof returns a formatted informational message with optional error.
func GetInfof(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)
	return GetInfo(text, opts)
}

// GetDebugf returns a formatted debug message with optional error.
func GetDebugf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)
	return GetDebug(text, opts)
}

// GetCriticalf returns a formatted critical message with optional error.
func GetCriticalf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)
	return GetCritical(text, opts)
}

// GetDeletef returns a formatted delete message with optional error.
func GetDeletef(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)
	return GetDelete(text, opts)
}

// GetUpdatef returns a formatted update message with optional error.
func GetUpdatef(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)
	return GetUpdate(text, opts)
}

// GetInstallf returns a formatted install message with optional error.
func GetInstallf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)
	return GetInstall(text, opts)
}

// GetUpgradef returns a formatted upgrade message with optional error.
func GetUpgradef(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)
	return GetUpgrade(text, opts)
}

// GetEditf returns a formatted edit message with optional error.
func GetEditf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)
	return GetEdit(text, opts)
}

// GetNewf returns a formatted creation message with optional error.
func GetNewf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)
	return GetNew(text, opts)
}

// GetDownloadf returns a formatted download message with optional error.
func GetDownloadf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)
	return GetDownload(text, opts)
}

// GetUploadf returns a formatted upload message with optional error.
func GetUploadf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)
	return GetUpload(text, opts)
}

// GetSyncf returns a formatted sync message with optional error.
func GetSyncf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)
	return GetSync(text, opts)
}

// GetSearchf returns a formatted search message with optional error.
func GetSearchf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)
	return GetSearch(text, opts)
}
