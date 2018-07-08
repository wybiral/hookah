package protocols

import (
	"net"
	"net/url"

	"github.com/wybiral/hookah/pkg/node"
)

// TCP creates a TCP client and returns Node
func TCP(addr string, opts url.Values) (node.Node, error) {
	rwc, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return rwc, nil
}
