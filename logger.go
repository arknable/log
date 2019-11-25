package log

import (
	"fmt"
	"io"
	golog "log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/arknable/errors"
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

const fileOutputExt = ".log"

// Options is configurable aspects of a Logger
type Options struct {
	// DisableStdOutput removes standard output
	DisableStdOutput bool

	// EnableFileOutput writes message to file
	EnableFileOutput bool

	// FileOutputFolder is path to folder where log files should be kept
	FileOutputFolder string

	// FileOutputName is the name of log file
	FileOutputName string
}

// Logger is a wrapper for Go's standard logger
type Logger struct {
	*golog.Logger
	Options

	lock               sync.Mutex
	currentFileOutName string
}

// New creates new logger
func New(opts Options) (*Logger, error) {
	l := &Logger{
		Options: opts,
	}
	writers, err := l.writers()
	if err != nil {
		return nil, errors.Wrap(err)
	}
	l.Logger = golog.New(writers, "", golog.LstdFlags)
	return l, nil
}

func (l *Logger) writers() (io.Writer, error) {
	w := make([]io.Writer, 0)
	if !l.DisableStdOutput {
		w = append(w, os.Stdout)
	}
	if l.EnableFileOutput {
		if len(l.FileOutputFolder) > 0 {
			if err := os.MkdirAll(l.FileOutputFolder, os.ModePerm); err != nil {
				return nil, errors.Wrap(err)
			}
		}
		l.currentFileOutName = fileName(l)
		filePath := filepath.Join(l.FileOutputFolder, l.currentFileOutName)
		file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, errors.Wrap(err)
		}
		w = append(w, file)
	}
	return io.MultiWriter(w...), nil
}

func fileName(l *Logger) string {
	return fmt.Sprintf("%s_%s%s", l.FileOutputName, time.Now().Format("20060102"), fileOutputExt)
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

// Logf mapped to Infof for compatibility with Segment's Logger
func (l *Logger) Logf(format string, v ...interface{}) {
	l.Infof(format, v...)
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
}
