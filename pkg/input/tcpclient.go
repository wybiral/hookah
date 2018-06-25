package input

import (
	"io"
	"net"
)

// Create a TCP client and return as ReadCloser
func tcpClient(addr string) (io.ReadCloser, error) {
	return net.Dial("tcp", addr)
}
