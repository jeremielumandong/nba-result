// Package utils provides shared utility functions used across the application.
package utils

import (
	"log"
	"os"
)

// Logger levels
const (
	LogLevelInfo  = "INFO"
	LogLevelWarn  = "WARN"
	LogLevelError = "ERROR"
	LogLevelDebug = "DEBUG"
)

var (
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
)

// init initializes loggers with appropriate prefixes and outputs.
func init() {
	infoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger = log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger = log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
}

// LogInfo logs an informational message.
func LogInfo(message string) {
	infoLogger.Println(message)
}

// LogWarn logs a warning message.
func LogWarn(message string) {
	warnLogger.Println(message)
}

// LogError logs an error message.
func LogError(message string) {
	errorLogger.Println(message)
}

// LogDebug logs a debug message.
func LogDebug(message string) {
	debugLogger.Println(message)
}

// LogInfof logs a formatted informational message.
func LogInfof(format string, args ...interface{}) {
	infoLogger.Printf(format, args...)
}

// LogWarnf logs a formatted warning message.
func LogWarnf(format string, args ...interface{}) {
	warnLogger.Printf(format, args...)
}

// LogErrorf logs a formatted error message.
func LogErrorf(format string, args ...interface{}) {
	errorLogger.Printf(format, args...)
}

// LogDebugf logs a formatted debug message.
func LogDebugf(format string, args ...interface{}) {
	debugLogger.Printf(format, args...)
}