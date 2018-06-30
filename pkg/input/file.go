package input

import (
	"io"
	"net/url"
	"os"
)

// File creates a file input and returns ReadCloser
func File(path string, args url.Values) (io.ReadCloser, error) {
	return os.Open(path)
}
