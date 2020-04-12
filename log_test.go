package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testLog(l *Logger) {
	l.Debug("This is a debug message")
	l.Info("This is", "an info message")
	l.Warning("This is a warning message")
	l.Error("This is", "an error message")

	l.Debugf("This is a %s message", "debugf")
	l.Infof("This is an %s message", "infof")
	l.Warningf("This is a %s message", "warningf")
	l.Errorf("This is an %s message", "errorf")
}

func TestNewDefault(t *testing.T) {
	assert.False(t, logger.DisableStdOut)

	err := NewDefault(&Options{
		DisableStdOut: true,
	})
	assert.Nil(t, err)
	assert.True(t, logger.DisableStdOut)
}

/* Manual tests */
// func TestFatalLevel(t *testing.T) {
// 	t.Log("Before log.Fatal")
// 	log.Fatal("This is fatal error")
// 	t.Log("After log.Fatal, this should not be diplayed")
// }
