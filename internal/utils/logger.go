package utils

import (
	"log"
	"os"
)

// LogLevel represents the severity of a log message
type LogLevel int

const (
	// Log levels
	DEBUG LogLevel = iota
	INFO
	ERROR
)

var (
	// Loggers for different levels
	infoLogger  *log.Logger
	debugLogger *log.Logger
	errorLogger *log.Logger

	// Current log level - can be changed at runtime
	currentLogLevel LogLevel = INFO
)

func init() {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
}

// SetLogLevel sets the minimum log level to display
func SetLogLevel(level LogLevel) {
	currentLogLevel = level
}

// Log logs a message at the specified level
func Log(level LogLevel, format string, v ...interface{}) {
	if level >= currentLogLevel {
		switch level {
		case DEBUG:
			debugLogger.Printf(format, v...)
		case INFO:
			infoLogger.Printf(format, v...)
		case ERROR:
			errorLogger.Printf(format, v...)
		}
	}
}

// Debug logs a debug message
func Debug(format string, v ...interface{}) {
	Log(DEBUG, format, v...)
}

// Info logs an info message
func Info(format string, v ...interface{}) {
	Log(INFO, format, v...)
}

// Error logs an error message
func Error(format string, v ...interface{}) {
	Log(ERROR, format, v...)
}
