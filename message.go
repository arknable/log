package log

import (
	"fmt"
	"os"

	"github.com/arknable/errors"
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
	msg := fmt.Sprintf(format, v...)
	l.Printf("%7s - %v\n", level.String(), msg)

}
