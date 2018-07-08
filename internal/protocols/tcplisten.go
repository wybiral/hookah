package protocols

import (
	"net/url"

	"github.com/wybiral/hookah/pkg/node"
)

// TCPListen creates a TCP server and returns Node
func TCPListen(addr string, opts url.Values) (node.Node, error) {
	return listen("tcp", addr)
}
