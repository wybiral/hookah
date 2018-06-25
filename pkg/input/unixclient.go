package input

import (
	"io"
	"net"
)

// Create a Unix client and return as ReadCloser
func unixClient(addr string) (io.ReadCloser, error) {
	return net.Dial("unix", addr)
}
