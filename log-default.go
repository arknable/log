package log

import (
	"github.com/arknable/errors"
)

// Default logger
var logger *Logger

// Default returns default logger
func Default() *Logger {
	return logger
}

// NewDefault creates new default logger
func NewDefault(opts *Options) error {
	l, err := New(opts)
	if err != nil {
		return errors.Wrap(err)
	}
	logger = l
	return nil
}
