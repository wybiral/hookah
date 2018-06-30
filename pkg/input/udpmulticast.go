package input

import (
	"io"
	"net"
	"net/url"
)

// UDPMulticast creates a UDP multicast listener and returns ReadCloser
func UDPMulticast(path string, args url.Values) (io.ReadCloser, error) {
	var err error
	var iface *net.Interface
	ifi := args.Get("iface")
	if len(ifi) > 0 {
		// If interface is supplied, look it up
		iface, err = net.InterfaceByName(ifi)
		if err != nil {
			return nil, err
		}
	}
	a, err := net.ResolveUDPAddr("udp", path)
	if err != nil {
		return nil, err
	}
	return net.ListenMulticastUDP("udp", iface, a)
}
