package protocols

import (
	"net/url"

	"github.com/wybiral/hookah/pkg/node"
)

// UnixListen creates a Unix domain socket listener and returns Node
func UnixListen(path string, opts url.Values) (node.Node, error) {
	return listen("unix", path)
}
