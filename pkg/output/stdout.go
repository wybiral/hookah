package output

import (
	"io"
	"os"
)

// Stdout returns stdout WriteCloser.
func Stdout(arg string) (io.WriteCloser, error) {
	return os.Stdout, nil
}

// Stderr returns stderr WriteCloser.
func Stderr(arg string) (io.WriteCloser, error) {
	return os.Stderr, nil
}
