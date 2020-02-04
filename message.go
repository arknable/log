package log

import (
	"fmt"
	"os"
	
	"github.com/arknable/errors"
	"github.com/fatih/color"
)

// Level is the level of this message
type Level string

// String returns string representation of this level
func (l Level) String() string {
	return (string) (l)
}

const (
	DebugLevel   = "DEBUG"
	InfoLevel    = "INFO"
	WarningLevel = "WARNING"
	ErrorLevel   = "ERROR"
	FatalLevel   = "FATAL"
)

var (
	italicStyle = color.New(color.Italic).SprintFunc()
	boldStyle   = color.New(color.Bold).SprintFunc()
	normalStyle = color.New(color.Reset).SprintFunc()

	messageFormat     = "%15s - %v"
	unformattedFormat = "%s\n"
)

// Prints log message with given format and level
func (l *Logger) write(level Level, format string, v ...interface{}) {
	var (
		msg       string
		msgFormat = messageFormat
	)

	if format == unformattedFormat {
		msgFormat = msgFormat + "\n"
		msg = fmt.Sprint(v...)
	} else {
		msg = fmt.Sprintf(format, v...)
	}

	if l.EnableFileOutput {
		fname := fileName(l)
		if fname != l.currentFileOutName {
			writers, err := l.writers()
			if err != nil {
				Fatal(errors.Wrap(err))
			}
			l.SetOutput(writers)
		}
	}
	
	levelString := normalStyle(level.String())
	if level == ErrorLevel || level == FatalLevel {
		levelString = boldStyle(level.String())
	} else if level == DebugLevel {
		levelString = italicStyle(level.String())
	}
	l.Printf(msgFormat, levelString, msg)

	if level == FatalLevel {
		os.Exit(1)
	}
}
