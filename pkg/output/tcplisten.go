package output

import (
	"io"
)

// Create a TCP server and return as ReadCloser
func tcpListen(addr string) (io.WriteCloser, error) {
	return listen("tcp", addr)
}
