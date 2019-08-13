package log

import (
	"fmt"
	"io"
	golog "log"
	"os"
	"sync"

	"github.com/arknable/errors"
	"github.com/fatih/color"
)

var (
	colorDebug   = color.New(color.Italic).SprintFunc()
	colorInfo    = color.New(color.Reset).SprintFunc()
	colorWarning = color.New(color.Bold).SprintFunc()
	colorError   = color.New(color.Bold).SprintFunc()
	colorFatal   = color.New(color.Bold).SprintFunc()

	messageFormat     = " -- %15s: %v"
	unformattedFormat = "%s\n"
)

type message struct {
	IsFormatted bool
	Format      string
	Level       string
	Message     string
}

// Logger is a wrapper for Go's standard logger
type Logger struct {
	*golog.Logger
	lock    sync.Mutex
	queue   []*message
	writers []io.Writer
}

// New creates new logger
func New() *Logger {
	return &Logger{
		Logger:  golog.New(os.Stdout, "", golog.LstdFlags),
		writers: make([]io.Writer, 0),
	}
}

// SetOutput sets output writer
func (l *Logger) SetOutput(w ...io.Writer) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.writers = w

	var writer io.Writer
	if len(l.writers) > 1 {
		writer = io.MultiWriter(l.writers...)
	} else if len(l.writers) == 1 {
		writer = l.writers[0]
	} else {
		writer = os.Stdout
	}
	l.Logger = golog.New(writer, "", golog.LstdFlags)
}

// AddFileOutput add output to file with given name
func (l *Logger) AddFileOutput(filePath string) (*os.File, error) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	l.writers = append(l.writers, file)
	l.SetOutput(l.writers...)
	return file, nil
}

// Hold puts log messages in a queueu, make sure to call Release() to write the messages
func (l *Logger) Hold() {
	l.queue = make([]*message, 0)
}

// Release writes queued messages.
// If something wrong happen during message writing,
// this function will make sure queue set to nil upon return.
func (l *Logger) Release() {
	defer func() {
		l.queue = nil
	}()
	for _, m := range l.queue {
		l.Printf(m.Format, m.Level, m.Message)
	}
}

// Debugf prints debug message with given format
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.printf(colorDebug("DEBUG"), format, v...)
}

// Debug prints debug message
func (l *Logger) Debug(v ...interface{}) {
	l.Debugf(unformattedFormat, v...)
}

// Infof prints info message with given format
func (l *Logger) Infof(format string, v ...interface{}) {
	l.printf(colorInfo("INFO"), format, v...)
}

// Info prints info message
func (l *Logger) Info(v ...interface{}) {
	l.Infof(unformattedFormat, v...)
}

// Warningf prints warning message with given format
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.printf(colorWarning("WARNING"), format, v...)
}

// Warning prints Warning message
func (l *Logger) Warning(v ...interface{}) {
	l.Warningf(unformattedFormat, v...)
}

// Errorf prints error message with given format
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.printf(colorError("ERROR"), format, v...)
}

// Error prints error message
func (l *Logger) Error(v ...interface{}) {
	l.Errorf(unformattedFormat, v...)
}

// Fatalf prints fatal message with given format
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.printf(colorFatal("FATAL"), format, v...)
}

// Fatal prints fatal message
func (l *Logger) Fatal(v ...interface{}) {
	l.Fatalf(unformattedFormat, v...)
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
	if l.queue != nil {
		l.queue = append(l.queue, msg)
		return
	}
	l.Printf(msg.Format, msg.Level, msg.Message)
}
