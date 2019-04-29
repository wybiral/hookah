package protocols

import (
	"net"
	"net/url"
	"strings"

	"github.com/wybiral/hookah/pkg/node"
)

// UDPListen creates a UDP listener Node
func UDPListen(addr string) (*node.Node, error) {
	a, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}
	c, err := net.ListenUDP("udp", a)
	if err != nil {
		return nil, err
	}
	n := node.New(c)
	n.W = nil
	return n, nil
}

// UDPMulticast creates a UDP multicast listener Node
func UDPMulticast(arg string) (*node.Node, error) {
	var err error
	var opts url.Values
	// Parse options
	addrOpts := strings.SplitN(arg, "?", 2)
	addr := addrOpts[0]
	if len(addrOpts) == 2 {
		op, err := url.ParseQuery(addrOpts[1])
		if err != nil {
			return nil, err
		}
		opts = op
	}
	iface := opts.Get("iface")
	var ifi *net.Interface
	ifi = nil
	if len(iface) > 0 {
		// If interface is supplied, look it up
		ifi, err = net.InterfaceByName(iface)
		if err != nil {
			return nil, err
		}
	}
	a, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}
	c, err := net.ListenMulticastUDP("udp", ifi, a)
	if err != nil {
		return nil, err
	}
	n := node.New(c)
	n.W = nil
	return n, nil
}
