package log

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFileWriter(t *testing.T) {
	w, err := NewFileWriter("test_new_writer", "./")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = w.Close(); err != nil {
			t.Fatal(err)
		}
		if err = os.Remove(w.filePath); err != nil {
			t.Fatal(err)
		}
	}()

	testText := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	n, err := w.Write([]byte(testText))
	assert.Nil(t, err)
	assert.Equal(t, len(testText), n)
}

func TestFileRotation(t *testing.T) {
	w, err := NewFileWriter("test_rotation", "./")
	if err != nil {
		t.Fatal(err)
	}
	w.RotationSize = 1024

	defer func() {
		if err = w.Close(); err != nil {
			t.Fatal(err)
		}

		name := strings.TrimSuffix(w.filePath, filepath.Ext(w.filePath))
		files, err := filepath.Glob(name + "*")
		if err != nil {
			t.Fatal(err)
		}
		for _, f := range files {
			if err = os.Remove(f); err != nil {
				t.Fatal(err)
			}
		}
	}()

	testText := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	for i := 0; i < 500; i++ {
		if _, err = w.Write([]byte(testText)); err != nil {
			t.Fatal(err)
		}
	}

	name := strings.TrimSuffix(w.filePath, filepath.Ext(w.filePath))
	files, err := filepath.Glob(name + "*")
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, len(files) > 1)
}
