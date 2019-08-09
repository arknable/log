package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/fatih/color"
)

var (
	colorDebug        = color.New(color.Italic).SprintFunc()
	colorInfo         = color.New(color.Reset).SprintFunc()
	colorWarning      = color.New(color.Bold).SprintFunc()
	colorError        = color.New(color.Bold).SprintFunc()
	colorFatal        = color.New(color.Bold).SprintFunc()
	messageFormat     = " -- %15s: %v"
	unformattedFormat = "%s\n"
	logger            *log.Logger
	writer            io.Writer
	lock              sync.Mutex
)

// Debugf prints debug message with given format
func Debugf(format string, v ...interface{}) {
	printf(colorDebug("DEBUG"), format, v...)
}

// Debug prints debug message
func Debug(v ...interface{}) {
	Debugf(unformattedFormat, v...)
}

// Infof prints info message with given format
func Infof(format string, v ...interface{}) {
	printf(colorInfo("INFO"), format, v...)
}

// Info prints info message
func Info(v ...interface{}) {
	Infof(unformattedFormat, v...)
}

// Warningf prints warning message with given format
func Warningf(format string, v ...interface{}) {
	printf(colorWarning("WARNING"), format, v...)
}

// Warning prints Warning message
func Warning(v ...interface{}) {
	Warningf(unformattedFormat, v...)
}

// Errorf prints error message with given format
func Errorf(format string, v ...interface{}) {
	printf(colorError("ERROR"), format, v...)
}

// Error prints error message
func Error(v ...interface{}) {
	Errorf(unformattedFormat, v...)
}

// Fatalf prints fatal message with given format
func Fatalf(format string, v ...interface{}) {
	printf(colorFatal("FATAL"), format, v...)
}

// Fatal prints fatal message
func Fatal(v ...interface{}) {
	Fatalf(unformattedFormat, v...)
}

// SetOutput sets output writer
func SetOutput(w io.Writer) {
	lock.Lock()
	defer lock.Unlock()
	writer = w
	if logger != nil {
		logger.SetOutput(writer)
	}
}

// Prints log message with given format and level
func printf(level, format string, v ...interface{}) {
	if logger == nil {
		logger = New()
	}
	if format != unformattedFormat {
		logger.Printf(messageFormat, level, fmt.Sprintf(format, v...))
	} else {
		logger.Printf(messageFormat+"\n", level, fmt.Sprint(v...))
	}
}

// New creates new logger
func New() *log.Logger {
	if writer == nil {
		writer = os.Stdout
	}
	return log.New(writer, "", log.LstdFlags)
}
