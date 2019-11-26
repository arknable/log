package log

import (
	"fmt"
	"os"

	"github.com/arknable/errors"
	"github.com/fatih/color"
)

var (
	colorDebug   = color.New(color.Italic).SprintFunc()
	colorInfo    = color.New(color.Reset).SprintFunc()
	colorWarning = color.New(color.Bold).SprintFunc()
	colorError   = color.New(color.Bold).SprintFunc()
	colorFatal   = color.New(color.Bold).SprintFunc()

	coloredFatalLevel = colorFatal(fatalLevel)

	messageFormat     = "%15s | %v"
	unformattedFormat = "%s\n"
)

type message struct {
	IsFormatted bool
	Format      string
	Level       string
	Message     string
}

// Prints log message with given format and level
func (l *Logger) printf(level, format string, v ...interface{}) {
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

	l.Printf(msgFormat, level, msg)

	if level == coloredFatalLevel {
		os.Exit(1)
	}
}
