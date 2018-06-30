package input

import (
	"io"
	"net/url"
	"os"
)

// Stdin returns stdin ReadCloser.
func Stdin(path string, args url.Values) (io.ReadCloser, error) {
	return os.Stdin, nil
}
