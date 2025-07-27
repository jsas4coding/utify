package main

import (
	"fmt"

	"github.com/jsas4coding/utify"
)

func main() {
	fmt.Println("=== Utify Icon Examples ===")

	// Check if Nerd Font was detected
	if utify.IsNerdFontDetected() {
		fmt.Println("🎉 Nerd Font detected! (Available for manual use)")
	} else {
		fmt.Println("📝 Nerd Font not detected. Using regular Unicode icons by default.")
	}

	fmt.Printf("Current icon type: %d (0=None, 1=Regular, 2=NerdFont)\n", utify.GetIconType())
	fmt.Println()
	fmt.Println("💡 Tip: Set NERD_FONT_ENABLED=true environment variable to force Nerd Font icons")
	fmt.Println("    Or use utify.ForceNerdFont() in your code")
	fmt.Println()

	// Test with icons enabled
	fmt.Println("--- With Icons Enabled ---")
	opts := utify.OptionsDefault().WithIcon()

	utify.Success("Operation completed successfully!", opts)
	utify.Error("An error occurred!", opts)
	utify.Warning("Pay attention to this!", opts)
	utify.Info("Here is some information.", opts)
	utify.Debug("Debug message for developers.", opts)
	utify.Critical("Critical system failure!", opts)

	fmt.Println()
	utify.Install("Package installed", opts)
	utify.Download("File downloaded", opts)
	utify.Upload("File uploaded", opts)
	utify.Update("System updated", opts)
	utify.Delete("File deleted", opts)
	utify.New("New item created", opts)

	fmt.Println("\n--- Without Icons ---")
	optsNoIcon := utify.OptionsDefault().WithoutIcon()

	utify.Success("Operation completed (no icon)", optsNoIcon)
	utify.Error("An error occurred (no icon)", optsNoIcon)
	utify.Warning("Pay attention (no icon)", optsNoIcon)

	fmt.Println("\n--- Force Different Icon Types ---")

	// Force regular icons
	fmt.Println("\nForcing Regular Unicode Icons:")
	utify.ForceRegularIcons()
	utify.Success("Regular icon success", opts)
	utify.Error("Regular icon error", opts)

	// Force Nerd Font icons
	fmt.Println("\nForcing Nerd Font Icons:")
	utify.ForceNerdFont()
	utify.Success("Nerd Font success", opts)
	utify.Error("Nerd Font error", opts)

	// Disable all icons
	fmt.Println("\nDisabling All Icons:")
	utify.DisableIcons()
	utify.Success("No icons at all", opts) // Won't show icons even with WithIcon()

	fmt.Println("\n--- Icon Detection Info ---")
	fmt.Printf("Nerd Font detected: %v\n", utify.IsNerdFontDetected())
	fmt.Printf("Current icon type: %d (0=None, 1=Regular, 2=NerdFont)\n", utify.GetIconType())
}
