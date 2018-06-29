package input

import (
	"io"
)

// Create a Unix listener and return as ReadCloser
func unixListen(addr string) (io.ReadCloser, error) {
	return listen("unix", addr)
}
