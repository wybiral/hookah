package output

import (
	"io"
)

// Create a Unix server and return as ReadCloser
func unixServer(addr string) (io.WriteCloser, error) {
	return listenServer("unix", addr)
}
