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
	return (string)(l)
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
)

// Prints log message with given format and level
func (l *Logger) write(level Level, format string, v ...interface{}) {
	if level == FatalLevel {
		os.Exit(1)
	}

	if l.EnableFileOutput {
		filename := fileName(l)
		if filename != l.currentFileOutName {
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

	msg := fmt.Sprintf(format, v...)
	l.Printf("%15s - %v\n", levelString, msg)

}
