package input

import (
	"io"
	"net"
)

// TCP creates a TCP client and return as ReadCloser
func TCP(addr string) (io.ReadCloser, error) {
	return net.Dial("tcp", addr)
}
