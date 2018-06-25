package output

import (
	"io"
	"net"
)

// Create a TCP client and return as WriteCloser
func tcpClient(addr string) (io.WriteCloser, error) {
	return net.Dial("tcp", addr)
}
