package utils

import (
	"os"
)

// HandleError logs an error and optionally exits the program
func HandleError(err error, exitOnError bool) {
	if err != nil {
		Error("Error: %v", err)
		if exitOnError {
			os.Exit(1)
		}
	}
}

// HandleErrorWithMessage logs an error with a custom message and optionally exits the program
func HandleErrorWithMessage(err error, message string, exitOnError bool) {
	if err != nil {
		Error("%s: %v", message, err)
		if exitOnError {
			os.Exit(1)
		}
	}
}

// Assert checks a condition and logs an error if it's false
func Assert(condition bool, message string, exitOnError bool) {
	if !condition {
		Error("Assertion failed: %s", message)
		if exitOnError {
			os.Exit(1)
		}
	}
}

// FatalError logs an error and exits the program
func FatalError(format string, v ...interface{}) {
	Error(format, v...)
	os.Exit(1)
}
