package input

import (
	"io"
)

// Create a TCP server and return as ReadCloser
func tcpServer(addr string) (io.ReadCloser, error) {
	return listenServer("tcp", addr)
}
