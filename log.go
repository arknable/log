package log

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

var (
	colorDebug        = color.New(color.FgBlack, color.Bold).SprintFunc()
	colorInfo         = color.New(color.FgBlack, color.Bold).SprintFunc()
	colorWarning      = color.New(color.FgBlack, color.Bold, color.BgHiYellow).SprintFunc()
	colorError        = color.New(color.FgBlack, color.Bold, color.BgHiRed).SprintFunc()
	colorFatal        = color.New(color.FgWhite, color.Bold, color.BgRed).SprintFunc()
	messageFormat     = "|%s| %v"
	unformattedFormat = "%s\n"
)

// Debugf prints debug message with given format
func Debugf(format string, v ...interface{}) {
	printf(colorDebug("\tDEBUG\t"), format, v...)
}

// Debug prints debug message
func Debug(v ...interface{}) {
	Debugf(unformattedFormat, v...)
}

// Infof prints info message with given format
func Infof(format string, v ...interface{}) {
	printf(colorInfo("\tINFO\t"), format, v...)
}

// Info prints info message
func Info(v ...interface{}) {
	Infof(unformattedFormat, v...)
}

// Warningf prints warning message with given format
func Warningf(format string, v ...interface{}) {
	printf(colorWarning("\tWARNING\t"), format, v...)
}

// Warning prints Warning message
func Warning(v ...interface{}) {
	Warningf(unformattedFormat, v...)
}

// Errorf prints error message with given format
func Errorf(format string, v ...interface{}) {
	printf(colorError("\tERROR\t"), format, v...)
}

// Error prints error message
func Error(v ...interface{}) {
	Errorf(unformattedFormat, v...)
}

// Fatalf prints fatal message with given format
func Fatalf(format string, v ...interface{}) {
	printf(colorFatal("\tFATAL\t"), format, v...)
}

// Fatal prints fatal message
func Fatal(v ...interface{}) {
	Fatalf(unformattedFormat, v...)
}

// Prints log message with given format and level
func printf(level, format string, v ...interface{}) {
	if format != unformattedFormat {
		log.Printf(messageFormat, level, fmt.Sprintf(format, v...))
	} else {
		log.Printf(messageFormat+"\n", level, fmt.Sprint(v...))
	}
}