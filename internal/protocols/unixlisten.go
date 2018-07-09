package protocols

import (
	"github.com/wybiral/hookah/pkg/node"
)

// UnixListen creates a Unix domain socket listener Node
func UnixListen(path string) (*node.Node, error) {
	return listen("unix", path)
}
