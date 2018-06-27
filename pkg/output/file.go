package output

import (
	"io"
	"os"
)

// Create a file output (in append mode) and return as WriteCloser
func file(path string) (io.WriteCloser, error) {
	flags := os.O_APPEND | os.O_WRONLY | os.O_CREATE
	return os.OpenFile(path, flags, 0600)
}
