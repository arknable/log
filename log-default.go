package log

// Default defaultLogger.
var defaultLogger = New()

// SetDefaultLogger replaces default logger with given one.
func SetDefaultLogger(l *Logger) {
	defaultLogger = l
}
