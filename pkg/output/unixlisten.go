package output

import (
	"io"
)

// UnixListen creates a Unix domain socket listener and returns ReadCloser
func UnixListen(addr string) (io.WriteCloser, error) {
	return listen("unix", addr)
}
