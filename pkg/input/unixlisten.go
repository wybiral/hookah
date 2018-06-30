package input

import (
	"io"
)

// UnixListen creates a Unix domain socket listener and returns ReadCloser
func UnixListen(addr string) (io.ReadCloser, error) {
	return listen("unix", addr)
}
