package main

import (
	"fmt"
	
	"github.com/jsas4coding/utify"
)

func main() {
	opts := utify.OptionsDefault()

	utify.Success("Operation completed successfully!", opts)
	utify.Error("An error occurred!", opts)
	utify.Warning("Pay attention to this!", opts)
	utify.Info("Here is some information.", opts)
	utify.Debug("Debug message for developers.", opts)
	utify.Critical("Critical system failure!", opts)

	// Using formatted versions
	utify.Successf("Processed %d items successfully", opts, 42)
	utify.Errorf("Failed to process item #%d", opts, 5)

	// Using styled options
	styledOpts := utify.OptionsDefault().
		WithBold().
		WithItalic()
	
	utify.Success("This message is bold and italic!", styledOpts)

	// Using icons
	fmt.Println("\n--- With Icons ---")
	iconOpts := utify.OptionsDefault().WithIcon()
	utify.Success("Success with icon!", iconOpts)
	utify.Error("Error with icon!", iconOpts)
	utify.Warning("Warning with icon!", iconOpts)

	// Using Get functions for manual output handling
	text, err := utify.GetError("Something went wrong", opts)
	if err == utify.ErrSilent {
		// Handle the error as needed
		_ = text
	}
}