package log

import (
	"fmt"
	stdlog "log"
	"os"
)

// Logger is a wrapper for Go's standard defaultLogger
type loggerImpl struct {
	stdlog.Logger
}

// New creates new defaultLogger
func New() Logger {
	l := &loggerImpl{
		Logger: stdlog.Logger{},
	}
	l.SetOutput(os.Stdout)
	l.SetFlags(stdlog.LstdFlags)
	return l
}

// Debugf prints debug message with given format
func (l *loggerImpl) Debugf(format string, v ...interface{}) {
	l.write(DebugLevel, format, v...)
}

// Debug prints debug message
func (l *loggerImpl) Debug(v ...interface{}) {
	l.writeln(DebugLevel, v...)
}

// Infof prints info message with given format
func (l *loggerImpl) Infof(format string, v ...interface{}) {
	l.write(InfoLevel, format, v...)
}

// Info prints info message
func (l *loggerImpl) Info(v ...interface{}) {
	l.writeln(InfoLevel, v...)
}

// Warningf prints warning message with given format
func (l *loggerImpl) Warningf(format string, v ...interface{}) {
	l.write(WarningLevel, format, v...)
}

// Warning prints Warning message
func (l *loggerImpl) Warning(v ...interface{}) {
	l.write(WarningLevel, "%v\n", v...)
}

// Errorf prints error message with given format
func (l *loggerImpl) Errorf(format string, v ...interface{}) {
	l.write(ErrorLevel, format, v...)
}

// Error prints error message
func (l *loggerImpl) Error(v ...interface{}) {
	l.writeln(ErrorLevel, v...)
}

// Fatalf prints fatal message with given format
func (l *loggerImpl) Fatalf(format string, v ...interface{}) {
	l.write(FatalLevel, format, v...)
}

// Fatal prints fatal message
func (l *loggerImpl) Fatal(v ...interface{}) {
	l.writeln(FatalLevel, v...)
}

// Logf mapped to Infof for compatibility with Segment's Logger
func (l *loggerImpl) Logf(format string, v ...interface{}) {
	l.Infof(format, v...)
}

// Prints log message with given format and level
func (l *loggerImpl) write(level Level, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.Printf("%7s %s", level.String(), msg)

	if level == FatalLevel {
		os.Exit(1)
	}
}

// Prints log message with given level
func (l *loggerImpl) writeln(level Level, v ...interface{}) {
	msg := []interface{}{
		fmt.Sprintf("%7s", level.String()),
	}
	l.Println(append(msg, v...)...)

	if level == FatalLevel {
		os.Exit(1)
	}
}
