package output

import (
	"io"
	"net"
)

// Create a UDP client and return as WriteCloser
func udpClient(addr string) (io.WriteCloser, error) {
	return net.Dial("udp", addr)
}
