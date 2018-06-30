package output

import (
	"io"
	"net/url"
	"os"
)

// File creates a file output (in append mode) and returns WriteCloser
func File(path string, opts url.Values) (io.WriteCloser, error) {
	flags := os.O_APPEND | os.O_WRONLY | os.O_CREATE
	return os.OpenFile(path, flags, 0600)
}
