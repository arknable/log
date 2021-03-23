package log

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDailyFileOutput(t *testing.T) {
	file, err := DailyFileOutput("test_daily_file", "./")
	if err != nil {
		t.Fatal(err)
	}

	filename := file.Name()
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	logger := New()
	logger.SetOutput(io.MultiWriter(os.Stdout, file))

	logger.Info("Test info log written to file")
	logger.Error("Test error log written to file")

	file.Close()

	_, err = os.Stat(file.Name())
	assert.Nil(t, err)

	file, err = os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	assert.Equal(t, 2, len(lines))
	assert.True(t, strings.HasSuffix(lines[0], "INFO Test info log written to file"))
	assert.True(t, strings.HasSuffix(lines[1], "ERROR Test error log written to file"))
}
