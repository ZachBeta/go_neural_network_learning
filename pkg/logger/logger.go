package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// LogLevel represents the severity of a log message
type LogLevel int

const (
	// DEBUG level for detailed debugging information
	DEBUG LogLevel = iota
	// INFO level for general information
	INFO
	// WARN level for warning messages
	WARN
	// ERROR level for error messages
	ERROR
)

// Logger is a wrapper around the standard log.Logger
type Logger struct {
	*log.Logger
	file     *os.File
	level    LogLevel
	mu       sync.Mutex
	filePath string
}

var (
	// Default logger instance
	defaultLogger *Logger
	// Log directory
	logDir = "logs"
)

// Init initializes the default logger
func Init() error {
	var err error
	defaultLogger, err = NewLogger("training", INFO)
	return err
}

// NewLogger creates a new logger with the given name and log level
func NewLogger(name string, level LogLevel) (*Logger, error) {
	// Create logs directory if it doesn't exist
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create logs directory: %w", err)
	}

	// Create a new log file with timestamp
	logFile := filepath.Join(logDir, fmt.Sprintf("%s_%s.log", name, time.Now().Format("2006-01-02_15-04-05")))
	f, err := os.Create(logFile)
	if err != nil {
		return nil, fmt.Errorf("failed to create log file: %w", err)
	}

	// Create a multi-writer to write to both file and stdout
	// For now, we'll only write to the file
	writer := io.Writer(f)

	// Create the logger
	logger := &Logger{
		Logger:   log.New(writer, "", log.LstdFlags),
		file:     f,
		level:    level,
		filePath: logFile,
	}

	return logger, nil
}

// Close closes the log file
func (l *Logger) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.file.Close()
}

// Debug logs a debug message
func (l *Logger) Debug(format string, v ...interface{}) {
	if l.level <= DEBUG {
		l.Printf("[DEBUG] "+format, v...)
	}
}

// Info logs an info message
func (l *Logger) Info(format string, v ...interface{}) {
	if l.level <= INFO {
		l.Printf("[INFO] "+format, v...)
	}
}

// Warn logs a warning message
func (l *Logger) Warn(format string, v ...interface{}) {
	if l.level <= WARN {
		l.Printf("[WARN] "+format, v...)
	}
}

// Error logs an error message
func (l *Logger) Error(format string, v ...interface{}) {
	if l.level <= ERROR {
		l.Printf("[ERROR] "+format, v...)
	}
}

// GetFilePath returns the path to the log file
func (l *Logger) GetFilePath() string {
	return l.filePath
}

// Debug logs a debug message to the default logger
func Debug(format string, v ...interface{}) {
	if defaultLogger != nil {
		defaultLogger.Debug(format, v...)
	}
}

// Info logs an info message to the default logger
func Info(format string, v ...interface{}) {
	if defaultLogger != nil {
		defaultLogger.Info(format, v...)
	}
}

// Warn logs a warning message to the default logger
func Warn(format string, v ...interface{}) {
	if defaultLogger != nil {
		defaultLogger.Warn(format, v...)
	}
}

// Error logs an error message to the default logger
func Error(format string, v ...interface{}) {
	if defaultLogger != nil {
		defaultLogger.Error(format, v...)
	}
}

// Close closes the default logger
func Close() error {
	if defaultLogger != nil {
		return defaultLogger.Close()
	}
	return nil
}
