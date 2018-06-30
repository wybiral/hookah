package output

import (
	"io"
	"net"
	"net/url"
)

// TCP creates a TCP client and returns WriteCloser
func TCP(addr string, args url.Values) (io.WriteCloser, error) {
	return net.Dial("tcp", addr)
}
