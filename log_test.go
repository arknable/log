package log

import (
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	Debug("This is a debug message")
	Info("This is", " an info message")
	Warning("This is a warning message")
	Error("This is", " an error message")
	Fatal("This is a fatal message")

	Debugf("This is a %s message", "debugf")
	Infof("This is an %s message", "infof")
	Warningf("This is a %s message", "warningf")
	Errorf("This is an %s message", "errorf")
	Fatalf("This is a %s message", "fatalf")
}

func TestLogFile(t *testing.T) {
	logFile, err := os.OpenFile("test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		logFile.Close()
		os.Remove("test.log")
	}()

	SetOutput(io.MultiWriter(os.Stdout, logFile))
	TestLog(t)
}

func TestLock(t *testing.T) {
	logFile, err := os.OpenFile("test-lock.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		logFile.Close()
		os.Remove("test-lock.log")
	}()
	go func() {
		for i := 1; i <= 10; i++ {
			Debug("Message A.", i)
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for i := 1; i <= 10; i++ {
			Debug("Message B.", i)
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(3 * time.Second)
	SetOutput(logFile)
	time.Sleep(3 * time.Second)
	SetOutput(os.Stdout)
	time.Sleep(30 * time.Second)
}

func TestHold(t *testing.T) {
	Debug("Before Hold() called")

	Hold()

	Debug("This is a debug message")
	Info("This is", " an info message")
	Warning("This is a warning message")
	Error("This is", " an error message")
	Fatal("This is a fatal message")

	Debugf("This is a %s message", "debugf")
	Infof("This is an %s message", "infof")
	Warningf("This is a %s message", "warningf")
	Errorf("This is an %s message", "errorf")
	Fatalf("This is a %s message", "fatalf")

	for i := 0; i < 5; i++ {
		fmt.Println("Waiting ...")
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Call Release()")
	Release()

	Debug("After Release() called")
}

func TestInstance(t *testing.T) {
	logFile, err := os.OpenFile("test-instance.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		logFile.Close()
		os.Remove("test-instance.log")
	}()

	logger := New()
	logger.SetOutput(io.MultiWriter(os.Stdout, logFile))

	logger.Debug("Debug from new instance")
	logger.Info("Info from new instance")
	logger.Warningf("Warningf %s new instance", "from")
	logger.Errorf("Errorf from %s", "new instance")
}
