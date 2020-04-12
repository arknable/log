package log

// Default defaultLogger.
var defaultLogger = New()

// SetDefaultLogger replaces default logger with given one.
func SetDefaultLogger(l *Logger) {
	defaultLogger = l
}

// Debugf is bridge for Logger.Debugf of default defaultLogger
func Debugf(format string, v ...interface{}) {
	defaultLogger.Debugf(format, v...)
}

// Debug is bridge for Logger.Debug of default defaultLogger
func Debug(v ...interface{}) {
	defaultLogger.Debug(v...)
}

// Infof is bridge for Logger.Infof of default defaultLogger
func Infof(format string, v ...interface{}) {
	defaultLogger.Infof(format, v...)
}

// Info is bridge for Logger.Info of default defaultLogger
func Info(v ...interface{}) {
	defaultLogger.Info(v...)
}

// Warningf is bridge for Logger.Warningf of default defaultLogger
func Warningf(format string, v ...interface{}) {
	defaultLogger.Warningf(format, v...)
}

// Warning is bridge for Logger.Warning of default defaultLogger
func Warning(v ...interface{}) {
	defaultLogger.Warning(v...)
}

// Errorf is bridge for Logger.Errorf of default defaultLogger
func Errorf(format string, v ...interface{}) {
	defaultLogger.Errorf(format, v...)
}

// Error is bridge for Logger.Error of default defaultLogger
func Error(v ...interface{}) {
	defaultLogger.Error(v...)
}

// Fatalf is bridge for Logger.Fatalf of default defaultLogger
func Fatalf(format string, v ...interface{}) {
	defaultLogger.Fatalf(format, v...)
}

// Fatal is bridge for Logger.Fatal of default defaultLogger
func Fatal(v ...interface{}) {
	defaultLogger.Fatal(v...)
}
