package log

import "testing"

func TestLogger(t *testing.T) {
	l := New()

	l.Debug("This is a debug message")
	l.Info("This is", "an info message")
	l.Warning("This is a warning message")
	l.Error("This is", "an error message")

	l.Debugf("This is a %s message", "debugf")
	l.Infof("This is an %s message", "infof")
	l.Warningf("This is a %s message", "warningf")
	l.Errorf("This is an %s message", "errorf")
}

/* Manual tests */
// func TestFatalLevel(t *testing.T) {
// 	t.Log("Before log.Fatal")
// 	log.Fatal("This is fatal error")
// 	t.Log("After log.Fatal, this should not be diplayed")
// }
