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
	// Build formatted message
	message := buildFormattedMessage(msgType, text, opts)

	// Output message and log
	fmt.Println(message)
	logger.LogMessage(msgType, text)

	// Handle callback or exit
	handleCallbackOrExit(msgType, text, opts)

	// Return appropriate result
	return handleReturnValue(msgType, text)
}

// buildFormattedMessage constructs the formatted message string
func buildFormattedMessage(msgType messages.Type, text string, opts *options.Options) string {
	color := getColorForMessage(msgType, opts)
	style := getStyleForMessage(opts)
	icon := getIconForMessage(msgType, opts)

	return fmt.Sprintf("%s%s%s%s%s", style, color, icon, text, colors.Reset)
}

// getColorForMessage returns the appropriate color based on options
func getColorForMessage(msgType messages.Type, opts *options.Options) string {
	if opts.NoColor {
		return ""
	}
	return messages.GetColor(msgType)
}

// getStyleForMessage returns the appropriate style based on options
func getStyleForMessage(opts *options.Options) string {
	if opts.NoStyle {
		return ""
	}
	style := ""
	if opts.Bold {
		style += colors.Bold
	}
	if opts.Italic {
		style += colors.Italic
	}
	return style
}

// getIconForMessage returns the appropriate icon based on options
func getIconForMessage(msgType messages.Type, opts *options.Options) string {
	if !opts.ShowIcons || opts.NoIcon {
		return ""
	}
	icon := icons.GetIcon(msgType)
	if icon != "" {
		icon += " " // Add space after icon
	}
	return icon
}

// handleCallbackOrExit handles callback execution or program exit
func handleCallbackOrExit(msgType messages.Type, text string, opts *options.Options) {
	if opts.Callback != nil {
		opts.Callback(msgType, text)
	} else if opts.Exit && messages.IsErrorType(msgType) {
		os.Exit(1)
	}
}

// handleReturnValue returns the appropriate value and error
func handleReturnValue(msgType messages.Type, text string) (string, error) {
	if messages.IsErrorType(msgType) {
		return text, ErrSilent
	}
	return text, nil
}
