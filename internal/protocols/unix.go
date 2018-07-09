package protocols

import (
	"net"

	"github.com/wybiral/hookah/pkg/node"
)

// Unix creates a Unix domain socket client Node
func Unix(addr string) (*node.Node, error) {
	rwc, err := net.Dial("unix", addr)
	if err != nil {
		return nil, err
	}
	return node.New(rwc), nil
}
