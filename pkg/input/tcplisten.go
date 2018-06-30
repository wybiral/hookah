package input

import (
	"io"
)

// TCPListen creates a TCP listener and returns ReadCloser
func TCPListen(addr string) (io.ReadCloser, error) {
	return listen("tcp", addr)
}
