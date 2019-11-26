package log

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	colorDebug   = color.New(color.Italic).SprintFunc()
	colorInfo    = color.New(color.Reset).SprintFunc()
	colorWarning = color.New(color.Bold).SprintFunc()
	colorError   = color.New(color.Bold).SprintFunc()
	colorFatal   = color.New(color.Bold).SprintFunc()

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
	var msg *message
	if format == unformattedFormat {
		msg = &message{
			Format:  messageFormat + "\n",
			Level:   level,
			Message: fmt.Sprint(v...),
		}
	} else {
		msg = &message{
			IsFormatted: true,
			Format:      messageFormat,
			Level:       level,
			Message:     fmt.Sprintf(format, v...),
		}
	}
	l.Printf(msg.Format, msg.Level, msg.Message)
	if level == colorFatal(fatalLevel) {
		os.Exit(1)
	}
}
