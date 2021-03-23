package log

import "io"

// Level is the level of this message
type Level string

// String returns string representation of this level
func (l Level) String() string {
	return (string)(l)
}

const (
	// DebugLevel informs information for debugging purpose
	DebugLevel = "DEBUG"

	// InfoLevel informs that there is a useful information
	InfoLevel = "INFO"

	// WarningLevel informs that we need to pay more attention on something
	WarningLevel = "WARNING"

	// ErrorLevel informs that an error occured
	ErrorLevel = "ERROR"

	// FatalLevel informs that we are having a panic
	FatalLevel = "FATAL"
)

// Logger is required specification for a logger
type Logger interface {
	// Debug logs debug messages
	Debug(v ...interface{})

	// Debugf logs formatted debug messages
	Debugf(format string, v ...interface{})

	// Info lofs info messages
	Info(v ...interface{})

	// Infof lofs formatted info messages
	Infof(format string, v ...interface{})

	// Warning logs warning messages
	Warning(v ...interface{})

	// Warningf logs formatted warning messages
	Warningf(format string, v ...interface{})

	// Error logs error messages
	Error(v ...interface{})

	// Errorf logs formatted error messages
	Errorf(format string, v ...interface{})

	// Fatal logs fatal messages
	Fatal(v ...interface{})

	// Fatalf logs formatted fatal messages
	Fatalf(format string, v ...interface{})

	// SetOutput sets output writer
	SetOutput(w io.Writer)
}
