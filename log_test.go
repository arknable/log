package log

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testLog(l *Logger) {
	l.Debug("This is a debug message")
	l.Info("This is", " an info message")
	l.Warning("This is a warning message")
	l.Error("This is", " an error message")

	l.Debugf("This is a %s message", "debugf")
	l.Infof("This is an %s message", "infof")
	l.Warningf("This is a %s message", "warningf")
	l.Errorf("This is an %s message", "errorf")
}

func TestFileOutput(t *testing.T) {
	l, err := New(Options{
		EnableFileOutput: true,
		FileOutputName:   "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		files, err := ioutil.ReadDir(".")
		if err != nil {
			t.Fatal(err)
		}
		for _, f := range files {
			if strings.HasPrefix(f.Name(), l.FileOutputName) {
				os.Remove(f.Name())
			}
		}
	}()
	testLog(l)
}

func TestFileOutputFolder(t *testing.T) {
	l, err := New(Options{
		EnableFileOutput: true,
		FileOutputFolder: "log",
		FileOutputName:   "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.RemoveAll(l.FileOutputFolder); err != nil {
			t.Fatal(err)
		}
	}()
	testLog(l)
}

func TestNewDefault(t *testing.T) {
	assert.False(t, logger.DisableStdOutput)
	assert.False(t, logger.EnableFileOutput)
	assert.Equal(t, "", logger.FileOutputFolder)
	assert.Equal(t, "", logger.FileOutputName)

	err := NewDefault(Options{
		DisableStdOutput: true,
		FileOutputFolder: "foolder",
		FileOutputName:   "foo",
	})
	assert.Nil(t, err)
	assert.True(t, logger.DisableStdOutput)
	assert.False(t, logger.EnableFileOutput)
	assert.Equal(t, "foolder", logger.FileOutputFolder)
	assert.Equal(t, "foo", logger.FileOutputName)
}

/* Manual tests */
// func TestFatalLevel(t *testing.T) {
// 	t.Log("Before log.Fatal")
// 	log.Fatal("This is fatal error")
// 	t.Log("After log.Fatal, this should not be diplayed")
// }

// func TestDailyFileChange(t *testing.T) {
// 	l, err := New(Options{
// 		EnableFileOutput: true,
// 		FileOutputFolder: "log",
// 		FileOutputName:   "test",
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	for {
// 		l.Debug("Test debug message")
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }
