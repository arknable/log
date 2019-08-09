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
	queue             []*message
)

type message struct {
	IsFormatted bool
	Format      string
	Level       string
	Message     string
}

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

// Hold puts log messages in a queueu, make sure to call Release() to write the messages
func Hold() {
	queue = make([]*message, 0)
}

// Release writes queued messages.
// If something wrong happen during message writing,
// this function will make sure queue set to nil upon return.
func Release() {
	defer func() {
		queue = nil
	}()
	for _, m := range queue {
		logger.Printf(m.Format, m.Level, m.Message)
	}
}

// Prints log message with given format and level
func printf(level, format string, v ...interface{}) {
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
	if queue != nil {
		queue = append(queue, msg)
		return
	}
	logger.Printf(msg.Format, msg.Level, msg.Message)
}

// New creates new logger
func New() *log.Logger {
	return log.New(writer, "", log.LstdFlags)
}

func init() {
	writer = os.Stdout
	logger = New()
}
