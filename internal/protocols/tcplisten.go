package protocols

import (
	"github.com/wybiral/hookah/pkg/node"
)

// TCPListen creates a TCP listener Node
func TCPListen(addr string) (*node.Node, error) {
	return listen("tcp", addr)
}
