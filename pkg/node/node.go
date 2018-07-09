package node

import (
	"io"
)

// Node is the main interface used by all hookah nodes.
type Node struct {
	R io.Reader
	W io.Writer
	C io.Closer
}

// New returns a new Node from a ReadWriteCloser.
func New(rwc io.ReadWriteCloser) *Node {
	return &Node{R: rwc, W: rwc, C: rwc}
}
