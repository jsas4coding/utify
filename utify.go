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

// Backward compatibility exports
type MessageType = messages.Type
type Options = options.Options

var (
	ErrSilent = formatter.ErrSilent
	Echo      = formatter.Echo

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

func OptionsDefault() *Options {
	return options.Default()
}

func SetColorTable(newColors map[string]string) {
	colors.SetColorTable(newColors)
}

// Logger configuration functions
func SetLogTarget(target string) error {
	return logger.SetLogTarget(target)
}

func GetLogTarget() string {
	return logger.GetLogTarget()
}

func SetLoggingEnabled(enabled bool) {
	logger.SetEnabled(enabled)
}

func IsLoggingEnabled() bool {
	return logger.IsEnabled()
}

func CloseLogger() {
	logger.Close()
}

// Icon configuration functions
func ForceNerdFont() {
	icons.ForceNerdFont()
}

func ForceRegularIcons() {
	icons.ForceRegularIcons()
}

func DisableIcons() {
	icons.DisableIcons()
}

func IsNerdFontDetected() bool {
	return icons.IsNerdFontDetected()
}

func GetIconType() int {
	return int(icons.GetIconType())
}

// Direct output functions
func Success(text string, opts *Options) {
	_, _ = Echo(MessageSuccess, text, opts)
}

func Error(text string, opts *Options) {
	_, _ = Echo(MessageError, text, opts)
}

func Warning(text string, opts *Options) {
	_, _ = Echo(MessageWarning, text, opts)
}

func Info(text string, opts *Options) {
	_, _ = Echo(MessageInfo, text, opts)
}

func Debug(text string, opts *Options) {
	_, _ = Echo(MessageDebug, text, opts)
}

func Critical(text string, opts *Options) {
	_, _ = Echo(MessageCritical, text, opts)
}

func Delete(text string, opts *Options) {
	_, _ = Echo(MessageDelete, text, opts)
}

func Update(text string, opts *Options) {
	_, _ = Echo(MessageUpdate, text, opts)
}

func Install(text string, opts *Options) {
	_, _ = Echo(MessageInstall, text, opts)
}

func Upgrade(text string, opts *Options) {
	_, _ = Echo(MessageUpgrade, text, opts)
}

func Edit(text string, opts *Options) {
	_, _ = Echo(MessageEdit, text, opts)
}

func New(text string, opts *Options) {
	_, _ = Echo(MessageNew, text, opts)
}

func Download(text string, opts *Options) {
	_, _ = Echo(MessageDownload, text, opts)
}

func Upload(text string, opts *Options) {
	_, _ = Echo(MessageUpload, text, opts)
}

func Sync(text string, opts *Options) {
	_, _ = Echo(MessageSync, text, opts)
}

func Search(text string, opts *Options) {
	_, _ = Echo(MessageSearch, text, opts)
}

// Formatted functions
func Successf(text string, opts *Options, args ...any) {
	Success(fmt.Sprintf(text, args...), opts)
}

func Errorf(text string, opts *Options, args ...any) {
	Error(fmt.Sprintf(text, args...), opts)
}

func Warningf(text string, opts *Options, args ...any) {
	Warning(fmt.Sprintf(text, args...), opts)
}

func Infof(text string, opts *Options, args ...any) {
	Info(fmt.Sprintf(text, args...), opts)
}

func Debugf(text string, opts *Options, args ...any) {
	Debug(fmt.Sprintf(text, args...), opts)
}

func Criticalf(text string, opts *Options, args ...any) {
	Critical(fmt.Sprintf(text, args...), opts)
}

func Deletef(text string, opts *Options, args ...any) {
	Delete(fmt.Sprintf(text, args...), opts)
}

func Updatef(text string, opts *Options, args ...any) {
	Update(fmt.Sprintf(text, args...), opts)
}

func Installf(text string, opts *Options, args ...any) {
	Install(fmt.Sprintf(text, args...), opts)
}

func Upgradef(text string, opts *Options, args ...any) {
	Upgrade(fmt.Sprintf(text, args...), opts)
}

func Editf(text string, opts *Options, args ...any) {
	Edit(fmt.Sprintf(text, args...), opts)
}

func Newf(text string, opts *Options, args ...any) {
	New(fmt.Sprintf(text, args...), opts)
}

func Downloadf(text string, opts *Options, args ...any) {
	Download(fmt.Sprintf(text, args...), opts)
}

func Uploadf(text string, opts *Options, args ...any) {
	Upload(fmt.Sprintf(text, args...), opts)
}

func Syncf(text string, opts *Options, args ...any) {
	Sync(fmt.Sprintf(text, args...), opts)
}

func Searchf(text string, opts *Options, args ...any) {
	Search(fmt.Sprintf(text, args...), opts)
}

// Get functions (return formatted output)
func GetSuccess(text string, opts *Options) (string, error) {
	return Echo(MessageSuccess, text, opts)
}

func GetError(text string, opts *Options) (string, error) {
	return Echo(MessageError, text, opts)
}

func GetWarning(text string, opts *Options) (string, error) {
	return Echo(MessageWarning, text, opts)
}

func GetInfo(text string, opts *Options) (string, error) {
	return Echo(MessageInfo, text, opts)
}

func GetDebug(text string, opts *Options) (string, error) {
	return Echo(MessageDebug, text, opts)
}

func GetCritical(text string, opts *Options) (string, error) {
	return Echo(MessageCritical, text, opts)
}

func GetDelete(text string, opts *Options) (string, error) {
	return Echo(MessageDelete, text, opts)
}

func GetUpdate(text string, opts *Options) (string, error) {
	return Echo(MessageUpdate, text, opts)
}

func GetInstall(text string, opts *Options) (string, error) {
	return Echo(MessageInstall, text, opts)
}

func GetUpgrade(text string, opts *Options) (string, error) {
	return Echo(MessageUpgrade, text, opts)
}

func GetEdit(text string, opts *Options) (string, error) {
	return Echo(MessageEdit, text, opts)
}

func GetNew(text string, opts *Options) (string, error) {
	return Echo(MessageNew, text, opts)
}

func GetDownload(text string, opts *Options) (string, error) {
	return Echo(MessageDownload, text, opts)
}

func GetUpload(text string, opts *Options) (string, error) {
	return Echo(MessageUpload, text, opts)
}

func GetSync(text string, opts *Options) (string, error) {
	return Echo(MessageSync, text, opts)
}

func GetSearch(text string, opts *Options) (string, error) {
	return Echo(MessageSearch, text, opts)
}

// Get formatted functions
func GetSuccessf(text string, opts *Options, args ...any) (string, error) {
	return GetSuccess(fmt.Sprintf(text, args...), opts)
}

func GetErrorf(text string, opts *Options, args ...any) (string, error) {
	return GetError(fmt.Sprintf(text, args...), opts)
}

func GetWarningf(text string, opts *Options, args ...any) (string, error) {
	return GetWarning(fmt.Sprintf(text, args...), opts)
}

func GetInfof(text string, opts *Options, args ...any) (string, error) {
	return GetInfo(fmt.Sprintf(text, args...), opts)
}

func GetDebugf(text string, opts *Options, args ...any) (string, error) {
	return GetDebug(fmt.Sprintf(text, args...), opts)
}

func GetCriticalf(text string, opts *Options, args ...any) (string, error) {
	return GetCritical(fmt.Sprintf(text, args...), opts)
}

func GetDeletef(text string, opts *Options, args ...any) (string, error) {
	return GetDelete(fmt.Sprintf(text, args...), opts)
}

func GetUpdatef(text string, opts *Options, args ...any) (string, error) {
	return GetUpdate(fmt.Sprintf(text, args...), opts)
}

func GetInstallf(text string, opts *Options, args ...any) (string, error) {
	return GetInstall(fmt.Sprintf(text, args...), opts)
}

func GetUpgradef(text string, opts *Options, args ...any) (string, error) {
	return GetUpgrade(fmt.Sprintf(text, args...), opts)
}

func GetEditf(text string, opts *Options, args ...any) (string, error) {
	return GetEdit(fmt.Sprintf(text, args...), opts)
}

func GetNewf(text string, opts *Options, args ...any) (string, error) {
	return GetNew(fmt.Sprintf(text, args...), opts)
}

func GetDownloadf(text string, opts *Options, args ...any) (string, error) {
	return GetDownload(fmt.Sprintf(text, args...), opts)
}

func GetUploadf(text string, opts *Options, args ...any) (string, error) {
	return GetUpload(fmt.Sprintf(text, args...), opts)
}

func GetSyncf(text string, opts *Options, args ...any) (string, error) {
	return GetSync(fmt.Sprintf(text, args...), opts)
}

func GetSearchf(text string, opts *Options, args ...any) (string, error) {
	return GetSearch(fmt.Sprintf(text, args...), opts)
}

// Log-only functions (don't print to stdout, only log)
func LogSuccess(text string) {
	logger.LogOnly(MessageSuccess, text)
}

func LogError(text string) {
	logger.LogOnly(MessageError, text)
}

func LogWarning(text string) {
	logger.LogOnly(MessageWarning, text)
}

func LogInfo(text string) {
	logger.LogOnly(MessageInfo, text)
}

func LogDebug(text string) {
	logger.LogOnly(MessageDebug, text)
}

func LogCritical(text string) {
	logger.LogOnly(MessageCritical, text)
}

func LogDelete(text string) {
	logger.LogOnly(MessageDelete, text)
}

func LogUpdate(text string) {
	logger.LogOnly(MessageUpdate, text)
}

func LogInstall(text string) {
	logger.LogOnly(MessageInstall, text)
}

func LogUpgrade(text string) {
	logger.LogOnly(MessageUpgrade, text)
}

func LogEdit(text string) {
	logger.LogOnly(MessageEdit, text)
}

func LogNew(text string) {
	logger.LogOnly(MessageNew, text)
}

func LogDownload(text string) {
	logger.LogOnly(MessageDownload, text)
}

func LogUpload(text string) {
	logger.LogOnly(MessageUpload, text)
}

func LogSync(text string) {
	logger.LogOnly(MessageSync, text)
}

func LogSearch(text string) {
	logger.LogOnly(MessageSearch, text)
}

// Log-only formatted functions
func LogSuccessf(text string, args ...any) {
	LogSuccess(fmt.Sprintf(text, args...))
}

func LogErrorf(text string, args ...any) {
	LogError(fmt.Sprintf(text, args...))
}

func LogWarningf(text string, args ...any) {
	LogWarning(fmt.Sprintf(text, args...))
}

func LogInfof(text string, args ...any) {
	LogInfo(fmt.Sprintf(text, args...))
}

func LogDebugf(text string, args ...any) {
	LogDebug(fmt.Sprintf(text, args...))
}

func LogCriticalf(text string, args ...any) {
	LogCritical(fmt.Sprintf(text, args...))
}

func LogDeletef(text string, args ...any) {
	LogDelete(fmt.Sprintf(text, args...))
}

func LogUpdatef(text string, args ...any) {
	LogUpdate(fmt.Sprintf(text, args...))
}

func LogInstallf(text string, args ...any) {
	LogInstall(fmt.Sprintf(text, args...))
}

func LogUpgradef(text string, args ...any) {
	LogUpgrade(fmt.Sprintf(text, args...))
}

func LogEditf(text string, args ...any) {
	LogEdit(fmt.Sprintf(text, args...))
}

func LogNewf(text string, args ...any) {
	LogNew(fmt.Sprintf(text, args...))
}

func LogDownloadf(text string, args ...any) {
	LogDownload(fmt.Sprintf(text, args...))
}

func LogUploadf(text string, args ...any) {
	LogUpload(fmt.Sprintf(text, args...))
}

func LogSyncf(text string, args ...any) {
	LogSync(fmt.Sprintf(text, args...))
}

func LogSearchf(text string, args ...any) {
	LogSearch(fmt.Sprintf(text, args...))
}