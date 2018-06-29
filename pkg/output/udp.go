package output

import (
	"io"
	"net"
)

// UDP creates a UDP client and return as WriteCloser
func UDP(addr string) (io.WriteCloser, error) {
	return net.Dial("udp", addr)
}
