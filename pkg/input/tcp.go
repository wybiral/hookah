package input

import (
	"io"
	"net"
	"net/url"
)

// TCP creates a TCP client and returns ReadCloser
func TCP(path string, args url.Values) (io.ReadCloser, error) {
	return net.Dial("tcp", path)
}
