package log

import (
	"fmt"
	"io"
	stdlog "log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/arknable/errors"
)

const (
	fileOutputExt = ".log"
	fatalLevel    = "FatalLevel"
)

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
	*stdlog.Logger
	Options

	lock               sync.Mutex
	currentFileOutName string
}

// New creates new logger
func New(opts *Options) (*Logger, error) {
	l := &Logger{}
	if opts != nil {
		l.Options = *opts
	} else {
		l.Options = Options{}
	}
	writers, err := l.writers()
	if err != nil {
		return nil, errors.Wrap(err)
	}
	l.Logger = stdlog.New(writers, "", stdlog.LstdFlags)
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
	l.write(DebugLevel, format, v...)
}

// Debug prints debug message
func (l *Logger) Debug(v ...interface{}) {
	l.Debugf(unformattedFormat, v...)
}

// Infof prints info message with given format
func (l *Logger) Infof(format string, v ...interface{}) {
	l.write(InfoLevel, format, v...)
}

// Info prints info message
func (l *Logger) Info(v ...interface{}) {
	l.Infof(unformattedFormat, v...)
}

// Warningf prints warning message with given format
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.write(WarningLevel, format, v...)
}

// Warning prints Warning message
func (l *Logger) Warning(v ...interface{}) {
	l.Warningf(unformattedFormat, v...)
}

// Errorf prints error message with given format
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.write(ErrorLevel, format, v...)
}

// Error prints error message
func (l *Logger) Error(v ...interface{}) {
	l.Errorf(unformattedFormat, v...)
}

// Fatalf prints fatal message with given format
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.write(FatalLevel, format, v...)
}

// Fatal prints fatal message
func (l *Logger) Fatal(v ...interface{}) {
	l.Fatalf(unformattedFormat, v...)
}

// Logf mapped to Infof for compatibility with Segment's Logger
func (l *Logger) Logf(format string, v ...interface{}) {
	l.Infof(format, v...)
}
