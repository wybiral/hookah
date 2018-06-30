package input

import (
	"io"
	"os"
)

// Stdin returns stdin ReadCloser.
func Stdin(arg string) (io.ReadCloser, error) {
	return os.Stdin, nil
}
