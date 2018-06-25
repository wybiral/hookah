package output

import (
	"io"
	"net"
)

// Create a Unix client and return as WriteCloser
func unixClient(addr string) (io.WriteCloser, error) {
	return net.Dial("unix", addr)
}
