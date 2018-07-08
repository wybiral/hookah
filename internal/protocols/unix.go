package protocols

import (
	"net"
	"net/url"

	"github.com/wybiral/hookah/pkg/node"
)

// Unix creates a Unix domain socket client and returns Node
func Unix(addr string, opts url.Values) (node.Node, error) {
	rwc, err := net.Dial("unix", addr)
	if err != nil {
		return nil, err
	}
	return rwc, nil
}
