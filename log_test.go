package log

import "testing"

func TestLog(t *testing.T) {
	Debug("This is a debug message")
	Info("This is an info message")
	Warning("This is a warning message")
	Error("This is an error message")
	Fatal("This is a fatal message")

	Debugf("This is a %s message", "debugf")
	Infof("This is an %s message", "infof")
	Warningf("This is a %s message", "warningf")
	Errorf("This is an %s message", "errorf")
	Fatalf("This is a %s message", "fatalf")
}
