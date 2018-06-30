package input

import (
	"io"
	"net"
	"strings"
)

// UDPMulticast creates a UDP multicast listener and returns ReadCloser
func UDPMulticast(arg string) (io.ReadCloser, error) {
	var addr string
	var err error
	var iface *net.Interface
	parts := strings.SplitN(arg, ",", 2)
	if len(parts) == 1 {
		addr = parts[0]
	} else {
		addr = parts[1]
		// If interface is supplied, look it up
		iface, err = net.InterfaceByName(parts[0])
		if err != nil {
			return nil, err
		}
	}
	a, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}
	return net.ListenMulticastUDP("udp", iface, a)
}
