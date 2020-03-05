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
	if err := l.checkFileOutput(); err != nil {
		l.Logger.Fatal(errors.Wrap(err))
	}

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
	if err := l.checkFileOutput(); err != nil {
		l.Logger.Fatal(errors.Wrap(err))
	}

	msg := []interface{}{
		header(level),
	}
	l.Println(append(msg, v...)...)

	if level == FatalLevel {
		os.Exit(1)
	}
}

func (l *Logger) checkFileOutput() error {
	if l.EnableFileOutput {
		filename := fileName(l)
		if filename != l.currentFileOutName {
			writers, err := l.writers()
			if err != nil {
				return errors.Wrap(err)
			}
			l.SetOutput(writers)
		}
	}
	return nil
}

func header(level Level) string {
	return fmt.Sprintf("%7s", level.String())
}
