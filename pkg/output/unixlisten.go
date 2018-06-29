package output

import (
	"io"
)

// Create a Unix server and return as ReadCloser
func unixListen(addr string) (io.WriteCloser, error) {
	return listen("unix", addr)
}
