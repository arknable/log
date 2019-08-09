package log

import (
	"io"
	"log"
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
		log.Fatal(err)
	}
	defer func() {
		logFile.Close()
		os.Remove("test.log")
	}()

	SetOutput(io.MultiWriter(os.Stdout, logFile))
	TestLog(t)
}

func TestLock(t *testing.T) {
	logFile, err := os.OpenFile("test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		logFile.Close()
		os.Remove("test.log")
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
