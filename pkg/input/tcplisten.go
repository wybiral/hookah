package input

import (
	"io"
)

// Create a TCP listener and return as ReadCloser
func tcpListen(addr string) (io.ReadCloser, error) {
	return listen("tcp", addr)
}
