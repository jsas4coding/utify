package unit

import (
	"os"
	"testing"

	"github.com/jsas4coding/utify/pkg/icons"
	"github.com/jsas4coding/utify/pkg/messages"
)

func TestIconDetection(t *testing.T) {
	detected := icons.IsNerdFontDetected()
	_ = detected
}

func TestSetIconType(t *testing.T) {
	original := icons.GetIconType()
	defer icons.SetIconType(original)

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
	original := icons.GetIconType()
	defer icons.SetIconType(original)

	icons.SetIconType(icons.NoIcons)
	icon := icons.GetIcon(messages.Success)
	if icon != "" {
		t.Errorf("Expected empty icon for NoIcons, got %q", icon)
	}

	icons.SetIconType(icons.RegularIcons)
	icon = icons.GetIcon(messages.Success)
	if icon == "" {
		t.Error("Expected non-empty icon for RegularIcons")
	}

	icons.SetIconType(icons.NerdFontIcons)
	icon = icons.GetIcon(messages.Success)
	if icon == "" {
		t.Error("Expected non-empty icon for NerdFontIcons")
	}
}

func TestIconHelperFunctions(t *testing.T) {
	original := icons.GetIconType()
	defer icons.SetIconType(original)

	icons.ForceNerdFont()
	if icons.GetIconType() != icons.NerdFontIcons {
		t.Error("ForceNerdFont should set NerdFontIcons")
	}

	icons.ForceRegularIcons()
	if icons.GetIconType() != icons.RegularIcons {
		t.Error("ForceRegularIcons should set RegularIcons")
	}

	icons.DisableIcons()
	if icons.GetIconType() != icons.NoIcons {
		t.Error("DisableIcons should set NoIcons")
	}
}

func TestAllMessageTypesHaveIcons(t *testing.T) {
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

	icons.SetIconType(icons.RegularIcons)
	for _, msgType := range messageTypes {
		icon := icons.GetIcon(msgType)
		if icon == "" {
			t.Errorf("Regular icon missing for message type: %s", msgType)
		}
	}

	icons.SetIconType(icons.NerdFontIcons)
	for _, msgType := range messageTypes {
		icon := icons.GetIcon(msgType)
		if icon == "" {
			t.Errorf("Nerd font icon missing for message type: %s", msgType)
		}
	}
}

func TestNerdFontDetectionEnvVar(t *testing.T) {
	original := os.Getenv("NERD_FONT_ENABLED")
	defer os.Setenv("NERD_FONT_ENABLED", original)

	os.Setenv("NERD_FONT_ENABLED", "true")
	icons.Init()
	if icons.GetIconType() != icons.NerdFontIcons {
		t.Error("NERD_FONT_ENABLED=true should force Nerd Font icons")
	}

	os.Setenv("NERD_FONT_ENABLED", "false")
	icons.Init()
	if icons.GetIconType() == icons.NerdFontIcons {
		t.Error("NERD_FONT_ENABLED=false should not force Nerd Font icons")
	}
}
