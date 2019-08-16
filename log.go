package log

import (
	"io"
	"os"
)

// Default logger
var logger *Logger

func init() {
	logger = New()
}

// DefaultLogger returns default logger
func DefaultLogger() *Logger {
	return logger
}

// Debugf is bridge for Logger.Debugf of default logger
func Debugf(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}

// Debug is bridge for Logger.Debug of default logger
func Debug(v ...interface{}) {
	logger.Debug(v...)
}

// Infof is bridge for Logger.Infof of default logger
func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

// Info is bridge for Logger.Info of default logger
func Info(v ...interface{}) {
	logger.Info(v...)
}

// Warningf is bridge for Logger.Warningf of default logger
func Warningf(format string, v ...interface{}) {
	logger.Warningf(format, v...)
}

// Warning is bridge for Logger.Warning of default logger
func Warning(v ...interface{}) {
	logger.Warning(v...)
}

// Errorf is bridge for Logger.Errorf of default logger
func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

// Error is bridge for Logger.Error of default logger
func Error(v ...interface{}) {
	logger.Error(v...)
}

// Fatalf is bridge for Logger.Fatalf of default logger
func Fatalf(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}

// Fatal is bridge for Logger.Fatal of default logger
func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

// SetOutput is bridge for Logger.SetOutput of default logger
func SetOutput(w ...io.Writer) {
	logger.SetOutput(w...)
}

// AddFileOutput is bridge for Logger.AddFileOutput of default logger
func AddFileOutput(filePath string) (*os.File, error) {
	return logger.AddFileOutput(filePath)
}
