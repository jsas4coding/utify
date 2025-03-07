package utify

import (
	"fmt"
	"log"
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
	ColorBold      = "\033[1m"
	ColorItalic    = "\033[3m"
	ColorReset     = "\033[0m"
)

type MessageType string

type Options struct {
	Bold    bool
	Italic  bool
	NoColor bool
	NoIcon  bool
	Exit    bool
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

var userColors = map[string]string{}

func SetColorTable(newColors map[string]string) {
	for key, value := range newColors {
		userColors[key] = value
	}
}

func getColor(msgType MessageType) string {
	if color, exists := userColors[string(msgType)]; exists {
		return color
	}

	return colors[msgType]
}

func Echo(msgType MessageType, text string, opts Options) {
	color := getColor(msgType)

	if opts.NoColor {
		color = ""
	}

	style := ""

	if opts.Bold {
		style += ColorBold
	}

	if opts.Italic {
		style += ColorItalic
	}

	message := fmt.Sprintf("%s%s %s%s", style, color, text, ColorReset)

	fmt.Println(message)

	log.Printf("[%s] %s", strings.ToUpper(string(msgType)), text)

	if opts.Exit {
		os.Exit(1)
	}
}

func Success(text string, opts Options)    { Echo(MessageSuccess, text, opts) }
func Error(text string, opts Options)      { Echo(MessageError, text, opts) }
func Warning(text string, opts Options)    { Echo(MessageWarning, text, opts) }
func Info(text string, opts Options)       { Echo(MessageInfo, text, opts) }
func Debug(text string, opts Options)      { Echo(MessageDebug, text, opts) }
func Search(text string, opts Options)     { Echo(MessageSearch, text, opts) }
func Sync(text string, opts Options)       { Echo(MessageSync, text, opts) }
func Download(text string, opts Options)   { Echo(MessageDownload, text, opts) }
func Refresh(text string, opts Options)    { Echo(MessageRefresh, text, opts) }
func Upload(text string, opts Options)     { Echo(MessageUpload, text, opts) }
func Delete(text string, opts Options)     { Echo(MessageDelete, text, opts) }
func Critical(text string, opts Options)   { Echo(MessageCritical, text, opts) }
func Git(text string, opts Options)        { Echo(MessageGit, text, opts) }
func New(text string, opts Options)        { Echo(MessageNew, text, opts) }
func Edit(text string, opts Options)       { Echo(MessageEdit, text, opts) }
func Update(text string, opts Options)     { Echo(MessageUpdate, text, opts) }
func Generation(text string, opts Options) { Echo(MessageGeneration, text, opts) }
func Find(text string, opts Options)       { Echo(MessageFind, text, opts) }
func Link(text string, opts Options)       { Echo(MessageLink, text, opts) }
func Unlink(text string, opts Options)     { Echo(MessageUnlink, text, opts) }
func Upgrade(text string, opts Options)    { Echo(MessageUpgrade, text, opts) }
func Install(text string, opts Options)    { Echo(MessageInstall, text, opts) }
func Font(text string, opts Options)       { Echo(MessageFont, text, opts) }
func Theme(text string, opts Options)      { Echo(MessageTheme, text, opts) }
func Icon(text string, opts Options)       { Echo(MessageIcon, text, opts) }
