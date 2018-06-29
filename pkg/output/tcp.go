package output

import (
	"io"
	"net"
)

// TCP creates a TCP client and return as WriteCloser
func TCP(addr string) (io.WriteCloser, error) {
	return net.Dial("tcp", addr)
}
