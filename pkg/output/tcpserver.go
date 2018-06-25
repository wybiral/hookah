package output

import (
	"io"
)

// Create a TCP server and return as ReadCloser
func tcpServer(addr string) (io.WriteCloser, error) {
	return listenServer("tcp", addr)
}
