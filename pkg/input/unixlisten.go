package input

import (
	"io"
)

// UnixListen creates a Unix domain socket listener and return as ReadCloser
func UnixListen(addr string) (io.ReadCloser, error) {
	return listen("unix", addr)
}
