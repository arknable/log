package log

import (
	"io"
	stdlog "log"
	"os"

	"github.com/arknable/errors"
)

// Options is configurable aspects of a Logger
type Options struct {
	// DisableStdOut removes standard output
	DisableStdOut bool
}

// Logger is a wrapper for Go's standard logger
type Logger struct {
	*stdlog.Logger
	Options
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
	if !l.DisableStdOut {
		w = append(w, os.Stdout)
	}
	return io.MultiWriter(w...), nil
}

// Debugf prints debug message with given format
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.write(DebugLevel, format, v...)
}

// Debug prints debug message
func (l *Logger) Debug(v ...interface{}) {
	l.writeln(DebugLevel, v...)
}

// Infof prints info message with given format
func (l *Logger) Infof(format string, v ...interface{}) {
	l.write(InfoLevel, format, v...)
}

// Info prints info message
func (l *Logger) Info(v ...interface{}) {
	l.writeln(InfoLevel, v...)
}

// Warningf prints warning message with given format
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.write(WarningLevel, format, v...)
}

// Warning prints Warning message
func (l *Logger) Warning(v ...interface{}) {
	l.writeln(WarningLevel, v...)
}

// Errorf prints error message with given format
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.write(ErrorLevel, format, v...)
}

// Error prints error message
func (l *Logger) Error(v ...interface{}) {
	l.writeln(ErrorLevel, v...)
}

// Fatalf prints fatal message with given format
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.write(FatalLevel, format, v...)
}

// Fatal prints fatal message
func (l *Logger) Fatal(v ...interface{}) {
	l.writeln(FatalLevel, v...)
}

// Logf mapped to Infof for compatibility with Segment's Logger
func (l *Logger) Logf(format string, v ...interface{}) {
	l.Infof(format, v...)
}
