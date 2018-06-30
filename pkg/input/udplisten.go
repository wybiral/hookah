package input

import (
	"io"
	"net"
	"net/url"
)

// UDPListen creates a UDP listener and returns ReadCloser
func UDPListen(path string, args url.Values) (io.ReadCloser, error) {
	a, err := net.ResolveUDPAddr("udp", path)
	if err != nil {
		return nil, err
	}
	return net.ListenUDP("udp", a)
}
