package log

import (
	"bufio"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
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
		if err = w.archive(); err != nil {
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

func (w *FileWriter) archive() error {
	files, err := filepath.Glob(fmt.Sprintf("%s*", filepath.Join(w.Directory, w.Name)))
	if err != nil {
		return err
	}
	fileName := fmt.Sprintf("%s.%d.gz", filepath.Join(w.Directory, w.Name), len(files))
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer close(file)

	writer := gzip.NewWriter(file)
	defer close(writer)

	buffer := make([]byte, 1<<10)
	if _, err = w.file.Seek(0, 0); err != nil {
		return err
	}
	if _, err = io.CopyBuffer(writer, bufio.NewReader(w.file), buffer); err != nil {
		return err
	}

	if err = w.file.Close(); err != nil {
		return err
	}
	if err = os.Remove(w.filePath); err != nil {
		return err
	}

	return w.createNewFile()
}

func (w *FileWriter) createNewFile() error {
	filePath := filepath.Join(w.Directory, fmt.Sprintf("%s.%s", w.Name, w.Extension))
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	w.filePath = filePath
	w.file = file
	return nil
}
