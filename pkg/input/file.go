package input

import (
	"io"
	"os"
)

// File creates a file input and returns ReadCloser
func File(path string) (io.ReadCloser, error) {
	return os.Open(path)
}
