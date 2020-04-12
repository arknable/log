package log

import (
	"fmt"
	"io"
	"os"
)

// To be used by defer
func close(c io.Closer) {
	if err := c.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
