package log

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

var (
	colorDebug   = color.New(color.FgBlack, color.Bold).SprintFunc()
	colorInfo    = color.New(color.FgBlack, color.Bold).SprintFunc()
	colorWarning = color.New(color.FgBlack, color.Bold, color.BgHiYellow).SprintFunc()
	colorError   = color.New(color.FgBlack, color.Bold, color.BgHiRed).SprintFunc()
	colorFatal   = color.New(color.FgWhite, color.Bold, color.BgRed).SprintFunc()
	messageFormat = "|%s| %v"
)

func Debugf(format string, v ...interface{}) {
	printf(colorDebug("\tDEBUG\t"), format, v...)
}

func Debug(v interface{}) {
	Debugf("%s\n", v)
}

func Infof(format string, v ...interface{}) {
	printf(colorInfo("\tINFO\t"), format, v...)
}

func Info(v interface{}) {
	Infof("%s\n", v)
}

func Warningf(format string, v ...interface{}) {
	printf(colorWarning("\tWARNING\t"), format, v...)
}

func Warning(v interface{}) {
	Warningf("%s\n", v)
}

func Errorf(format string, v ...interface{}) {
	printf(colorError("\tERROR\t"), format, v...)
}

func Error(v interface{}) {
	Errorf("%s\n", v)
}

func Fatalf(format string, v ...interface{}) {
	printf(colorFatal("\tFATAL\t"), format, v...)
}

func Fatal(v interface{}) {
	Fatalf("%s\n", v)
}

func printf(level, format string, v ...interface{}) {
	log.Printf(messageFormat, level, fmt.Sprintf(format, v...))
}