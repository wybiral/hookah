package input

import (
	"io"
)

// Create a Unix server and return as ReadCloser
func unixServer(addr string) (io.ReadCloser, error) {
	return listenServer("unix", addr)
}
