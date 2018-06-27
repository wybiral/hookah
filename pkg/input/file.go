package input

import (
	"io"
	"os"
)

// Create a file input and return as ReadCloser
func file(path string) (io.ReadCloser, error) {
	return os.Open(path)
}
