package main

import (
	"fmt"

	"github.com/jsas4coding/utify"
)

func main() {
	// Define a callback function for logging or metrics
	callback := func(msgType utify.MessageType, message string) {
		fmt.Printf("📝 Callback: [%s] %s\n", msgType, message)
	}

	opts := utify.OptionsDefault().WithCallback(callback)

	utify.Success("Operation completed!", opts)
	utify.Error("Something failed!", opts)
	utify.Warning("Be careful!", opts)

	// The callback will be triggered for each message above
}
