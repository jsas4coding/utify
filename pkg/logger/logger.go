package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jsas4coding/utify/pkg/messages"
)

type LogEntry struct {
	Timestamp string        `json:"timestamp"`
	Level     string        `json:"level"`
	Message   string        `json:"message"`
	Type      messages.Type `json:"type"`
	Binary    string        `json:"binary"`
}

var (
	logFile   *os.File
	logger    *log.Logger
	logTarget string
	binaryName string
	enabled   = true
)

func init() {
	binaryName = getBinaryName()
	logTarget = fmt.Sprintf("/var/log/%s.log", binaryName)
	initLogger()
}

func getBinaryName() string {
	if len(os.Args) > 0 {
		return filepath.Base(os.Args[0])
	}
	return "utify"
}

func initLogger() {
	if !enabled {
		return
	}

	// Try to create log directory if it doesn't exist
	logDir := filepath.Dir(logTarget)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		// If we can't create the directory, fall back to current directory
		logTarget = fmt.Sprintf("%s.log", binaryName)
		logDir = "."
	}

	var err error
	logFile, err = os.OpenFile(logTarget, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		// Try fallback location in current directory
		fallbackTarget := fmt.Sprintf("%s.log", binaryName)
		logFile, err = os.OpenFile(fallbackTarget, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			// If we still can't open a log file, disable logging
			enabled = false
			return
		}
		logTarget = fallbackTarget
	}

	logger = log.New(logFile, "", 0)
}

func SetLogTarget(target string) error {
	if logFile != nil {
		logFile.Close()
	}
	
	logTarget = target
	initLogger()
	
	if !enabled {
		return fmt.Errorf("failed to initialize logger with target: %s", target)
	}
	
	return nil
}

func GetLogTarget() string {
	return logTarget
}

func SetEnabled(enable bool) {
	enabled = enable
	if !enabled && logFile != nil {
		logFile.Close()
		logFile = nil
		logger = nil
	} else if enabled && logFile == nil {
		initLogger()
	}
}

func IsEnabled() bool {
	return enabled
}

func LogMessage(msgType messages.Type, message string) {
	if !enabled || logger == nil {
		return
	}

	entry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Level:     strings.ToUpper(string(msgType)),
		Message:   message,
		Type:      msgType,
		Binary:    binaryName,
	}

	jsonData, err := json.Marshal(entry)
	if err != nil {
		// Fallback to simple format if JSON fails
		logger.Printf("[%s] %s", entry.Level, message)
		return
	}

	logger.Println(string(jsonData))
}

func LogOnly(msgType messages.Type, message string) {
	LogMessage(msgType, message)
}

func Close() {
	if logFile != nil {
		logFile.Close()
		logFile = nil
		logger = nil
	}
}