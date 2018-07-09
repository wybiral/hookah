package app

import "github.com/wybiral/hookah/pkg/node"

// Update contains the outgoing bytes and source node
type Update struct {
	// Bytes being sent
	Bytes []byte
	// Source of update
	Source *node.Node
}
