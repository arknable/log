package log

import (
	"fmt"
	"os"
)

// Level is the level of this message
type Level string

// String returns string representation of this level
func (l Level) String() string {
	return (string)(l)
}

const (
	// DebugLevel informs information for debugging purpose
	DebugLevel = "DBG"

	// InfoLevel informs that there is a useful information
	InfoLevel = "INF"

	// WarningLevel informs that we need to pay more attention on something
	WarningLevel = "WRN"

	// ErrorLevel informs that an error occured
	ErrorLevel = "ERR"

	// FatalLevel informs that we are having a panic
	FatalLevel = "FTL"
)

// Prints log message with given format and level
func (l *Logger) write(level Level, format string, v ...interface{}) {
	msg := []interface{}{
		header(level),
		fmt.Sprintf(format, v...),
	}
	l.Println(msg...)

	if level == FatalLevel {
		os.Exit(1)
	}
}

func (l *Logger) writeln(level Level, v ...interface{}) {
	msg := []interface{}{
		header(level),
	}
	l.Println(append(msg, v...)...)

	if level == FatalLevel {
		os.Exit(1)
	}
}

func header(level Level) string {
	return fmt.Sprintf("[%3s]", level.String())
}
