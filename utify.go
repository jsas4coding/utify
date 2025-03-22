package utify

import (
	"errors"
	"fmt"
	"log"
	"maps"
	"os"
	"strings"
)

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

type MessageType string

type Options struct {
	Bold     bool
	Italic   bool
	NoColor  bool
	NoIcon   bool
	NoStyle  bool
	Exit     bool
	Callback func(MessageType, string)
}

func OptionsDefault() *Options {
	return &Options{}
}

func (o *Options) WithBold() *Options {
	o.Bold = true
	return o
}

func (o *Options) WithItalic() *Options {
	o.Italic = true
	return o
}

func (o *Options) WithoutColor() *Options {
	o.NoColor = true
	return o
}

func (o *Options) WithoutIcon() *Options {
	o.NoIcon = true
	return o
}

func (o *Options) WithoutStyle() *Options {
	o.NoStyle = true
	return o
}

func (o *Options) WithExit() *Options {
	o.Exit = true
	o.Callback = nil
	return o
}

func (o *Options) WithCallback(cb func(MessageType, string)) *Options {
	o.Callback = cb
	o.Exit = false
	return o
}

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

var ErrSilent = errors.New("silent error")
var userColors = map[string]string{}

func SetColorTable(newColors map[string]string) {
	maps.Copy(userColors, newColors)
}

func getColor(msgType MessageType) string {
	if color, exists := userColors[string(msgType)]; exists {
		return color
	}

	return colors[msgType]
}

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

func Success(text string, opts *Options) (string, error) {
	Echo(MessageSuccess, text, opts)
	return text, nil
}

func Error(text string, opts *Options) (string, error) {
	Echo(MessageError, text, opts)
	return text, ErrSilent
}

func Warning(text string, opts *Options) (string, error) {
	Echo(MessageWarning, text, opts)
	return text, nil
}

func Info(text string, opts *Options) (string, error) {
	Echo(MessageInfo, text, opts)
	return text, nil
}
func Debug(text string, opts *Options) (string, error) {
	Echo(MessageDebug, text, opts)
	return text, nil
}

func Critical(text string, opts *Options) (string, error) {
	Echo(MessageCritical, text, opts)
	return text, ErrSilent
}

func Delete(text string, opts *Options) (string, error) {
	Echo(MessageDelete, text, opts)
	return text, nil
}

func Update(text string, opts *Options) (string, error) {
	Echo(MessageUpdate, text, opts)
	return text, nil
}

func Install(text string, opts *Options) (string, error) {
	Echo(MessageInstall, text, opts)
	return text, nil
}

func Upgrade(text string, opts *Options) (string, error) {
	Echo(MessageUpgrade, text, opts)
	return text, nil
}

func Edit(text string, opts *Options) (string, error) {
	Echo(MessageEdit, text, opts)
	return text, nil
}

func New(text string, opts *Options) (string, error) {
	Echo(MessageNew, text, opts)
	return text, nil
}

func Download(text string, opts *Options) (string, error) {
	Echo(MessageDownload, text, opts)
	return text, nil
}

func Upload(text string, opts *Options) (string, error) {
	Echo(MessageUpload, text, opts)
	return text, nil
}

func Sync(text string, opts *Options) (string, error) {
	Echo(MessageSync, text, opts)
	return text, nil
}

func Search(text string, opts *Options) (string, error) {
	Echo(MessageSearch, text, opts)
	return text, nil
}

func Successf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)

	return Success(text, opts)
}

func Errorf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)

	return Error(text, opts)
}

func Warningf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)

	return Warning(text, opts)
}

func Infof(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)

	return Info(text, opts)
}

func Debugf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)

	return Debug(text, opts)
}

func Criticalf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)

	return Critical(text, opts)
}

func Deletef(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)

	return Delete(text, opts)
}

func Updatef(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)

	return Update(text, opts)
}

func Installf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)

	return Install(text, opts)
}

func Upgradef(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)

	return Upgrade(text, opts)
}

func Editf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)

	return Edit(text, opts)
}

func Newf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)

	return New(text, opts)
}

func Downloadf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)

	return Download(text, opts)
}

func Uploadf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)

	return Upload(text, opts)
}

func Syncf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)

	return Sync(text, opts)
}

func Searchf(text string, opts *Options, args ...any) (string, error) {
	text = fmt.Sprintf(text, args...)

	return Search(text, opts)
}
