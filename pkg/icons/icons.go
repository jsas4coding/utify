package icons

import (
	"os"
	"strings"

	"github.com/jsas4coding/utify/pkg/messages"
)

// IconType represents different icon sets
type IconType int

const (
	NoIcons IconType = iota
	RegularIcons
	NerdFontIcons
)

// Nerd Font icons (requires Nerd Font) - using more commonly available codepoints
var nerdFontIcons = map[messages.Type]string{
	messages.Success:    "\uf00c", // nf-fa-check (Font Awesome check)
	messages.Error:      "\uf00d", // nf-fa-times (Font Awesome times)
	messages.Warning:    "\uf071", // nf-fa-exclamation_triangle (Font Awesome)
	messages.Info:       "\uf129", // nf-fa-info (Font Awesome info)
	messages.Debug:      "\uf188", // nf-fa-bug (Font Awesome bug)
	messages.Critical:   "\uf06a", // nf-fa-exclamation_circle (Font Awesome)
	messages.Search:     "\uf002", // nf-fa-search (Font Awesome search)
	messages.Sync:       "\uf021", // nf-fa-refresh (Font Awesome refresh)
	messages.Download:   "\uf019", // nf-fa-download (Font Awesome download)
	messages.Refresh:    "\uf021", // nf-fa-refresh (Font Awesome refresh)
	messages.Upload:     "\uf093", // nf-fa-upload (Font Awesome upload)
	messages.Delete:     "\uf1f8", // nf-fa-trash (Font Awesome trash)
	messages.Git:        "\ue702", // nf-dev-git (Devicons git)
	messages.New:        "\uf067", // nf-fa-plus (Font Awesome plus)
	messages.Edit:       "\uf040", // nf-fa-pencil (Font Awesome pencil)
	messages.Update:     "\uf021", // nf-fa-refresh (Font Awesome refresh)
	messages.Generation: "\uf013", // nf-fa-cog (Font Awesome cog)
	messages.Find:       "\uf002", // nf-fa-search (Font Awesome search)
	messages.Link:       "\uf0c1", // nf-fa-link (Font Awesome link)
	messages.Unlink:     "\uf127", // nf-fa-unlink (Font Awesome unlink)
	messages.Upgrade:    "\uf062", // nf-fa-arrow_up (Font Awesome arrow up)
	messages.Install:    "\uf019", // nf-fa-download (Font Awesome download)
	messages.Font:       "\uf031", // nf-fa-font (Font Awesome font)
	messages.Theme:      "\uf1fc", // nf-fa-paint_brush (Font Awesome paint brush)
	messages.Icon:       "\uf118", // nf-fa-smile_o (Font Awesome smile)
	messages.Default:    "\uf111", // nf-fa-circle (Font Awesome circle)
}

// Regular Unicode icons (fallback)
var regularIcons = map[messages.Type]string{
	messages.Success:    "✅", // check mark
	messages.Error:      "❌", // cross mark
	messages.Warning:    "⚠️ ", // warning sign
	messages.Info:       "ℹ️ ", // information
	messages.Debug:      "🐛", // bug
	messages.Critical:   "🚨", // rotating light
	messages.Search:     "🔍", // magnifying glass
	messages.Sync:       "🔄", // arrows counterclockwise
	messages.Download:   "⬇️ ", // down arrow
	messages.Refresh:    "🔃", // clockwise arrows
	messages.Upload:     "⬆️ ", // up arrow
	messages.Delete:     "🗑️ ", // wastebasket
	messages.Git:        "📦", // package
	messages.New:        "➕", // plus sign
	messages.Edit:       "✏️ ", // pencil
	messages.Update:     "🔄", // arrows counterclockwise
	messages.Generation: "⚙️ ", // gear
	messages.Find:       "🔎", // magnifying glass tilted right
	messages.Link:       "🔗", // link
	messages.Unlink:     "⛓️‍💥", // broken chain
	messages.Upgrade:    "⬆️ ", // up arrow
	messages.Install:    "📥", // inbox tray
	messages.Font:       "🔤", // latin letters
	messages.Theme:      "🎨", // artist palette
	messages.Icon:       "😀", // grinning face
	messages.Default:    "●",  // bullet
}

var currentIconType IconType
var detectedNerdFont bool

func init() {
	detectedNerdFont = detectNerdFont()
	// Check if user explicitly wants Nerd Font icons
	if os.Getenv("NERD_FONT_ENABLED") == "true" || os.Getenv("NERD_FONT_ENABLED") == "1" {
		currentIconType = NerdFontIcons
	} else {
		// Default to regular icons for safety - users can enable Nerd Fonts manually
		// This ensures compatibility across different terminal environments
		currentIconType = RegularIcons
	}
}

// detectNerdFont checks if a Nerd Font is likely available
func detectNerdFont() bool {
	// First check for explicit user preference
	if nerdFontEnv := os.Getenv("NERD_FONT_ENABLED"); nerdFontEnv != "" {
		return strings.ToLower(nerdFontEnv) == "true" || nerdFontEnv == "1"
	}
	
	// Check environment variables commonly set by terminals with Nerd Font support
	termProgram := os.Getenv("TERM_PROGRAM")
	terminal := os.Getenv("TERMINAL")
	term := os.Getenv("TERM")
	
	// Check for common terminals that typically use Nerd Fonts
	nerdFontTerminals := []string{
		"alacritty", "kitty", "wezterm", "hyper", "rio",
		"ghostty", "konsole", "gnome-terminal", "tilix",
		"terminator", "iterm", "warp", "tabby",
	}
	
	termProgram = strings.ToLower(termProgram)
	terminal = strings.ToLower(terminal)
	term = strings.ToLower(term)
	
	for _, termName := range nerdFontTerminals {
		if strings.Contains(termProgram, termName) || 
		   strings.Contains(terminal, termName) ||
		   strings.Contains(term, termName) {
			return true
		}
	}
	
	// Check font-related environment variables
	if font := os.Getenv("FONT"); font != "" {
		font = strings.ToLower(font)
		if strings.Contains(font, "nerd") || strings.Contains(font, "nerdfont") {
			return true
		}
	}
	
	// Additional checks for specific terminal features
	// Some terminals set specific environment variables
	if os.Getenv("KITTY_WINDOW_ID") != "" || 
	   os.Getenv("ALACRITTY_SOCKET") != "" ||
	   os.Getenv("WEZTERM_EXECUTABLE") != "" ||
	   os.Getenv("ITERM_SESSION_ID") != "" {
		return true
	}
	
	return false
}

// GetIcon returns the appropriate icon for a message type
func GetIcon(msgType messages.Type) string {
	switch currentIconType {
	case NerdFontIcons:
		if icon, exists := nerdFontIcons[msgType]; exists {
			return icon
		}
		return nerdFontIcons[messages.Default]
	case RegularIcons:
		if icon, exists := regularIcons[msgType]; exists {
			return icon
		}
		return regularIcons[messages.Default]
	default:
		return ""
	}
}

// SetIconType manually sets the icon type
func SetIconType(iconType IconType) {
	currentIconType = iconType
}

// GetIconType returns the current icon type
func GetIconType() IconType {
	return currentIconType
}

// IsNerdFontDetected returns whether Nerd Font was auto-detected
func IsNerdFontDetected() bool {
	return detectedNerdFont
}

// ForceNerdFont forces the use of Nerd Font icons
func ForceNerdFont() {
	SetIconType(NerdFontIcons)
}

// ForceRegularIcons forces the use of regular Unicode icons
func ForceRegularIcons() {
	SetIconType(RegularIcons)
}

// DisableIcons disables all icons
func DisableIcons() {
	SetIconType(NoIcons)
}