package formatter

import (
	"errors"
	"fmt"
	"os"

	"github.com/jsas4coding/utify/pkg/colors"
	"github.com/jsas4coding/utify/pkg/icons"
	"github.com/jsas4coding/utify/pkg/logger"
	"github.com/jsas4coding/utify/pkg/messages"
	"github.com/jsas4coding/utify/pkg/options"
)

var ErrSilent = errors.New("silent error")

func Echo(msgType messages.Type, text string, opts *options.Options) (string, error) {
	color := messages.GetColor(msgType)

	if opts.NoColor {
		color = ""
	}

	style := ""
	if opts.Bold {
		style += colors.Bold
	}
	if opts.Italic {
		style += colors.Italic
	}
	if opts.NoStyle {
		style = ""
	}

	// Handle icon logic
	icon := ""
	if opts.ShowIcons && !opts.NoIcon {
		icon = icons.GetIcon(msgType)
		if icon != "" {
			icon += " " // Add space after icon
		}
	}

	message := fmt.Sprintf("%s%s%s%s%s", style, color, icon, text, colors.Reset)

	fmt.Println(message)
	logger.LogMessage(msgType, text)

	if opts.Callback != nil {
		opts.Callback(msgType, text)
	} else if opts.Exit && messages.IsErrorType(msgType) {
		os.Exit(1)
	}

	if messages.IsErrorType(msgType) {
		return text, ErrSilent
	}

	return text, nil
}
