package output

import (
	"io"
	"net/url"
	"os"
)

// Stdout returns stdout WriteCloser.
func Stdout(path string, opts url.Values) (io.WriteCloser, error) {
	return os.Stdout, nil
}

// Stderr returns stderr WriteCloser.
func Stderr(path string, opts url.Values) (io.WriteCloser, error) {
	return os.Stderr, nil
}
