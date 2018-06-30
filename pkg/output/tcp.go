package output

import (
	"io"
	"net"
	"net/url"
)

// TCP creates a TCP client and returns WriteCloser
func TCP(path string, args url.Values) (io.WriteCloser, error) {
	return net.Dial("tcp", path)
}
