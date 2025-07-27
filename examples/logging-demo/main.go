package main

import (
	"fmt"

	"github.com/jsas4coding/utify"
)

func main() {
	fmt.Println("=== Utify Logging Examples ===")

	// Check current log target
	fmt.Printf("Current log target: %s\n", utify.GetLogTarget())
	fmt.Printf("Logging enabled: %v\n", utify.IsLoggingEnabled())

	// Normal functions (print to stdout AND log to file)
	utify.Success("This goes to both stdout and log file", utify.OptionsDefault())
	utify.Error("This error goes to both stdout and log file", utify.OptionsDefault())

	// Log-only functions (only write to log file, no stdout output)
	utify.LogInfo("This only goes to the log file")
	utify.LogWarning("This warning only goes to the log file")

	// Log-only formatted functions
	utify.LogSuccessf("Successfully processed %d items", 42)
	utify.LogErrorf("Failed to process item #%d: %s", 5, "connection timeout")

	// Change log target
	customLogPath := "./custom_app.log"
	err := utify.SetLogTarget(customLogPath)
	if err != nil {
		fmt.Printf("Failed to set custom log target: %v\n", err)
	} else {
		fmt.Printf("Changed log target to: %s\n", customLogPath)
		utify.LogInfo("This message goes to the custom log file")
	}

	// Temporarily disable logging
	utify.SetLoggingEnabled(false)
	utify.LogInfo("This message will NOT be logged")
	fmt.Println("Logging disabled - the above message was not logged")

	// Re-enable logging
	utify.SetLoggingEnabled(true)
	utify.LogInfo("Logging re-enabled - this message will be logged")

	// Clean up
	utify.CloseLogger()
	fmt.Printf("\nCheck the log files at:\n- %s\n- %s\n", utify.GetLogTarget(), customLogPath)
}