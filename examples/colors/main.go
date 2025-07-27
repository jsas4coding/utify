package main

import "github.com/jsas4coding/utify"

func main() {
	// Set custom colors
	utify.SetColorTable(map[string]string{
		string(utify.MessageSuccess): "\033[95m", // Purple
		string(utify.MessageError):   "\033[96m", // Light cyan
		string(utify.MessageWarning): "\033[93m", // Bright yellow
	})

	opts := utify.OptionsDefault()

	utify.Success("This success message is now purple!", opts)
	utify.Error("This error message is now light cyan!", opts)
	utify.Warning("This warning message is now bright yellow!", opts)
	utify.Info("This info message uses the default cyan color", opts)
}