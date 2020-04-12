package log

import (
	stdlog "log"
	"os"
)

// Logger is a wrapper for Go's standard defaultLogger
type Logger struct {
	stdlog.Logger
}

// New creates new defaultLogger
func New() *Logger {
	l := &Logger{
		Logger: stdlog.Logger{},
	}
	l.SetOutput(os.Stdout)
	l.SetFlags(stdlog.LstdFlags)
	return l
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
