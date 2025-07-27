package unit

import (
	"testing"

	"github.com/jsas4coding/utify/pkg/icons"
	"github.com/jsas4coding/utify/pkg/messages"
)

func TestIconDetection(t *testing.T) {
	// Test that icon detection doesn't crash
	detected := icons.IsNerdFontDetected()
	_ = detected // Just ensure it returns a bool
}

func TestSetIconType(t *testing.T) {
	// Save original
	original := icons.GetIconType()
	defer icons.SetIconType(original)

	// Test setting different icon types
	icons.SetIconType(icons.NoIcons)
	if icons.GetIconType() != icons.NoIcons {
		t.Error("Expected NoIcons type")
	}

	icons.SetIconType(icons.RegularIcons)
	if icons.GetIconType() != icons.RegularIcons {
		t.Error("Expected RegularIcons type")
	}

	icons.SetIconType(icons.NerdFontIcons)
	if icons.GetIconType() != icons.NerdFontIcons {
		t.Error("Expected NerdFontIcons type")
	}
}

func TestGetIcon(t *testing.T) {
	// Save original
	original := icons.GetIconType()
	defer icons.SetIconType(original)

	// Test NoIcons
	icons.SetIconType(icons.NoIcons)
	icon := icons.GetIcon(messages.Success)
	if icon != "" {
		t.Errorf("Expected empty icon for NoIcons, got %q", icon)
	}

	// Test RegularIcons
	icons.SetIconType(icons.RegularIcons)
	icon = icons.GetIcon(messages.Success)
	if icon == "" {
		t.Error("Expected non-empty icon for RegularIcons")
	}
	if icon != "✅" {
		t.Errorf("Expected ✅ for success, got %q", icon)
	}

	// Test NerdFontIcons
	icons.SetIconType(icons.NerdFontIcons)
	icon = icons.GetIcon(messages.Success)
	if icon == "" {
		t.Error("Expected non-empty icon for NerdFontIcons")
	}
}

func TestIconHelperFunctions(t *testing.T) {
	// Save original
	original := icons.GetIconType()
	defer icons.SetIconType(original)

	// Test ForceNerdFont
	icons.ForceNerdFont()
	if icons.GetIconType() != icons.NerdFontIcons {
		t.Error("ForceNerdFont should set NerdFontIcons")
	}

	// Test ForceRegularIcons
	icons.ForceRegularIcons()
	if icons.GetIconType() != icons.RegularIcons {
		t.Error("ForceRegularIcons should set RegularIcons")
	}

	// Test DisableIcons
	icons.DisableIcons()
	if icons.GetIconType() != icons.NoIcons {
		t.Error("DisableIcons should set NoIcons")
	}
}

func TestAllMessageTypesHaveIcons(t *testing.T) {
	// Save original
	original := icons.GetIconType()
	defer icons.SetIconType(original)

	messageTypes := []messages.Type{
		messages.Success, messages.Error, messages.Warning, messages.Info,
		messages.Debug, messages.Critical, messages.Search, messages.Sync,
		messages.Download, messages.Refresh, messages.Upload, messages.Delete,
		messages.Git, messages.New, messages.Edit, messages.Update,
		messages.Generation, messages.Find, messages.Link, messages.Unlink,
		messages.Upgrade, messages.Install, messages.Font, messages.Theme,
		messages.Icon, messages.Default,
	}

	// Test Regular Icons
	icons.SetIconType(icons.RegularIcons)
	for _, msgType := range messageTypes {
		icon := icons.GetIcon(msgType)
		if icon == "" {
			t.Errorf("Regular icon missing for message type: %s", msgType)
		}
	}

	// Test Nerd Font Icons
	icons.SetIconType(icons.NerdFontIcons)
	for _, msgType := range messageTypes {
		icon := icons.GetIcon(msgType)
		if icon == "" {
			t.Errorf("Nerd font icon missing for message type: %s", msgType)
		}
	}
}
