package protocols

import (
	"net"

	"github.com/wybiral/hookah/pkg/node"
)

// TCP creates a TCP client node
func TCP(addr string) (*node.Node, error) {
	rwc, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return node.New(rwc), nil
}
