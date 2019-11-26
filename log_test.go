package log

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func testLog(l *Logger) {
	l.Debug("This is a debug message")
	l.Info("This is", " an info message")
	l.Warning("This is a warning message")
	l.Error("This is", " an error message")
	l.Fatal("This is a fatal message")

	l.Debugf("This is a %s message", "debugf")
	l.Infof("This is an %s message", "infof")
	l.Warningf("This is a %s message", "warningf")
	l.Errorf("This is an %s message", "errorf")
	l.Fatalf("This is a %s message", "fatalf")
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
