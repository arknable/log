package log

import "testing"

func TestLogging(t *testing.T) {
	Debug("This is a debug message")
	Info("This is", "an info message")
	Warning("This is a warning message")
	Error("This is", "an error message")

	Debugf("This is a %s message", "debugf")
	Infof("This is an %s message", "infof")
	Warningf("This is a %s message", "warningf")
	Errorf("This is an %s message", "errorf")
}

/* Manual tests */
// func TestFatalLevel(t *testing.T) {
// 	t.Log("Before log.Fatal")
// 	log.Fatal("This is fatal error")
// 	t.Log("After log.Fatal, this should not be diplayed")
// }
