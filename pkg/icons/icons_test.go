package icons

import (
	"os"
	"testing"

	"github.com/jsas4coding/utify/pkg/messages"
)

func TestIconDetection(t *testing.T) {
	detected := IsNerdFontDetected()
	_ = detected
}

func TestSetIconType(t *testing.T) {
	original := GetIconType()
	defer SetIconType(original)

	SetIconType(NoIcons)
	if GetIconType() != NoIcons {
		t.Error("Expected NoIcons type")
	}

	SetIconType(RegularIcons)
	if GetIconType() != RegularIcons {
		t.Error("Expected RegularIcons type")
	}

	SetIconType(NerdFontIcons)
	if GetIconType() != NerdFontIcons {
		t.Error("Expected NerdFontIcons type")
	}
}

func TestGetIcon(t *testing.T) {
	original := GetIconType()
	defer SetIconType(original)

	SetIconType(NoIcons)
	icon := GetIcon(messages.Success)
	if icon != "" {
		t.Errorf("Expected empty icon for NoIcons, got %q", icon)
	}

	SetIconType(RegularIcons)
	icon = GetIcon(messages.Success)
	if icon == "" {
		t.Error("Expected non-empty icon for RegularIcons")
	}

	SetIconType(NerdFontIcons)
	icon = GetIcon(messages.Success)
	if icon == "" {
		t.Error("Expected non-empty icon for NerdFontIcons")
	}
}

func TestIconHelperFunctions(t *testing.T) {
	original := GetIconType()
	defer SetIconType(original)

	ForceNerdFont()
	if GetIconType() != NerdFontIcons {
		t.Error("ForceNerdFont should set NerdFontIcons")
	}

	ForceRegularIcons()
	if GetIconType() != RegularIcons {
		t.Error("ForceRegularIcons should set RegularIcons")
	}

	DisableIcons()
	if GetIconType() != NoIcons {
		t.Error("DisableIcons should set NoIcons")
	}
}

func TestAllMessageTypesHaveIcons(t *testing.T) {
	original := GetIconType()
	defer SetIconType(original)

	messageTypes := []messages.Type{
		messages.Success, messages.Error, messages.Warning, messages.Info,
		messages.Debug, messages.Critical, messages.Search, messages.Sync,
		messages.Download, messages.Refresh, messages.Upload, messages.Delete,
		messages.Git, messages.New, messages.Edit, messages.Update,
		messages.Generation, messages.Find, messages.Link, messages.Unlink,
		messages.Upgrade, messages.Install, messages.Font, messages.Theme,
		messages.Icon, messages.Default,
	}

	SetIconType(RegularIcons)
	for _, msgType := range messageTypes {
		icon := GetIcon(msgType)
		if icon == "" {
			t.Errorf("Regular icon missing for message type: %s", msgType)
		}
	}

	SetIconType(NerdFontIcons)
	for _, msgType := range messageTypes {
		icon := GetIcon(msgType)
		if icon == "" {
			t.Errorf("Nerd font icon missing for message type: %s", msgType)
		}
	}
}

func TestNerdFontDetectionEnvVar(t *testing.T) {
	original := os.Getenv("NERD_FONT_ENABLED")
	defer func() {
		if err := os.Setenv("NERD_FONT_ENABLED", original); err != nil {
			t.Errorf("Failed to restore original NERD_FONT_ENABLED value: %v", err)
		}
	}()

	if err := os.Setenv("NERD_FONT_ENABLED", "true"); err != nil {
		t.Errorf("Failed to set NERD_FONT_ENABLED=true: %v", err)
	}
	Init()
	if GetIconType() != NerdFontIcons {
		t.Error("NERD_FONT_ENABLED=true should force Nerd Font icons")
	}

	if err := os.Setenv("NERD_FONT_ENABLED", "false"); err != nil {
		t.Errorf("Failed to set NERD_FONT_ENABLED=false: %v", err)
	}
	Init()
	if GetIconType() == NerdFontIcons {
		t.Error("NERD_FONT_ENABLED=false should not force Nerd Font icons")
	}
}
