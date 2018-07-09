package protocols

import (
	"os"

	"github.com/wybiral/hookah/pkg/node"
)

// Stdin creates a stdin Node
func Stdin(path string) (*node.Node, error) {
	return &node.Node{R: os.Stdin}, nil
}

// Stdout creates a stdout Node
func Stdout(path string) (*node.Node, error) {
	return &node.Node{W: os.Stdout}, nil
}

// Stderr creates a stderr Node
func Stderr(path string) (*node.Node, error) {
	return &node.Node{W: os.Stderr}, nil
}

// Stdio creates a stdio Node
func Stdio(path string) (*node.Node, error) {
	return &node.Node{R: os.Stdin, W: os.Stdout}, nil
}
