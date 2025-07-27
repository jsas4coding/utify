package messages

import "github.com/jsas4coding/utify/pkg/colors"

type Type string

const (
	Success    Type = "success"
	Error      Type = "error"
	Warning    Type = "warning"
	Info       Type = "info"
	Debug      Type = "debug"
	Search     Type = "search"
	Sync       Type = "sync"
	Download   Type = "download"
	Refresh    Type = "refresh"
	Upload     Type = "upload"
	Delete     Type = "delete"
	Critical   Type = "critical"
	Git        Type = "git"
	New        Type = "new"
	Edit       Type = "edit"
	Update     Type = "update"
	Generation Type = "generation"
	Find       Type = "find"
	Link       Type = "link"
	Unlink     Type = "unlink"
	Upgrade    Type = "upgrade"
	Install    Type = "install"
	Font       Type = "font"
	Theme      Type = "theme"
	Icon       Type = "icon"
	Default    Type = "default"
)

var defaultColors = map[Type]string{
	Success:    colors.Green,
	Error:      colors.Red,
	Warning:    colors.Yellow,
	Info:       colors.Cyan,
	Debug:      colors.Gray,
	Search:     colors.Blue,
	Sync:       colors.Magenta,
	Download:   colors.White,
	Refresh:    colors.LightBlue,
	Upload:     colors.Green,
	Delete:     colors.Red,
	Critical:   colors.Magenta,
	Git:        colors.Magenta,
	New:        colors.Green,
	Edit:       colors.Blue,
	Update:     colors.Yellow,
	Generation: colors.Cyan,
	Find:       colors.Blue,
	Link:       colors.Magenta,
	Unlink:     colors.Red,
	Upgrade:    colors.LightBlue,
	Install:    colors.Green,
	Font:       colors.White,
	Theme:      colors.Magenta,
	Icon:       colors.White,
	Default:    colors.White,
}

func GetColor(msgType Type) string {
	if color, exists := colors.GetUserColor(string(msgType)); exists {
		return color
	}
	return defaultColors[msgType]
}

func IsErrorType(msgType Type) bool {
	return msgType == Error || msgType == Critical || msgType == Debug
}
