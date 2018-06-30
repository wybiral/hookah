package output

import (
	"io"
	"net"
	"net/url"
)

// UDP creates a UDP client and returns WriteCloser
func UDP(path string, args url.Values) (io.WriteCloser, error) {
	return net.Dial("udp", path)
}
