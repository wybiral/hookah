package input

import (
	"io"
	"net"
	"net/url"
)

// TCP creates a TCP client and returns ReadCloser
func TCP(addr string, opts url.Values) (io.ReadCloser, error) {
	return net.Dial("tcp", addr)
}
