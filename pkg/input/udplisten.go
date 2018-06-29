package input

import (
	"io"
	"net"
)

// Create a UDP listener and return as ReadCloser
func udpListen(addr string) (io.ReadCloser, error) {
	a, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}
	return net.ListenUDP("udp", a)
}
