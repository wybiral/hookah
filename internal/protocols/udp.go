package protocols

import (
	"net"

	"github.com/wybiral/hookah/pkg/node"
)

// UDP creates a UDP client node
func UDP(addr string) (*node.Node, error) {
	rwc, err := net.Dial("udp", addr)
	if err != nil {
		return nil, err
	}
	return node.New(rwc), nil
}
