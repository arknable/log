package log

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// ErrFileNotOpened occurs when log
var ErrFileNotOpened = errors.New("log file not yet opened for writing")

// FileWriter writes log messages into a text file.
// The log file will be rotated, i.e. replaced with new one,
// if the size larger than RotationSize.
type FileWriter struct {
	// Name is name of the log file.
	Name string

	// Directory is full path of the folder in which log files will be written.
	// This writer assumes that the directory is already created with proper permission.
	Directory string

	// Extension is the file extension of the log file.
	Extension string

	// RotationSize is the maximum file size allowed before rotation takes place.
	// Default size is 1 MB.
	RotationSize int64

	file     *os.File
	filePath string
}

// NewFileWriter creates new file writer
func NewFileWriter(name, folder string) (*FileWriter, error) {
	w := &FileWriter{
		Name:         name,
		Directory:    folder,
		Extension:    "log",
		RotationSize: int64(1 << 20),
	}

	if err := w.createNewFile(); err != nil {
		return nil, err
	}
	return w, nil
}

// Write implements io.Writer
func (w *FileWriter) Write(p []byte) (int, error) {
	if w.file == nil {
		return 0, ErrFileNotOpened
	}

	inf, err := w.file.Stat()
	if err != nil {
		return 0, err
	}
	if inf.Size() > w.RotationSize {
		if err = w.file.Sync(); err != nil {
			return 0, err
		}
		if err = w.file.Close(); err != nil {
			return 0, err
		}

		if err = os.Rename(w.filePath, fmt.Sprintf("%s.%s", w.filePath, time.Now().Format("2006Jan2_150405"))); err != nil {
			return 0, err
		}
		if err := w.createNewFile(); err != nil {
			return 0, err
		}
	}

	return w.file.Write(p)
}

// Close implements io.Closer
func (w *FileWriter) Close() error {
	if w.file != nil {
		return w.file.Close()
	}
	return nil
}

func (w *FileWriter) createNewFile() error {
	w.filePath = filepath.Join(w.Directory, fmt.Sprintf("%s.%s", w.Name, w.Extension))
	file, err := os.OpenFile(w.filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	w.file = file
	return nil
}
