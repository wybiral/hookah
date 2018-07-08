package protocols

import (
	"net/url"
	"os"

	"github.com/wybiral/hookah/pkg/node"
)

// Stdin creates a stdin Node
func Stdin(path string, opts url.Values) (node.Node, error) {
	return os.Stdin, nil
}

// Stdout creates a stdout Node
func Stdout(path string, opts url.Values) (node.Node, error) {
	return os.Stdout, nil
}

// Stderr creates a stderr Node
func Stderr(path string, opts url.Values) (node.Node, error) {
	return os.Stderr, nil
}

// Stdio creates a stdio Node
func Stdio(path string, opts url.Values) (node.Node, error) {
	return stdio{}, nil
}

type stdio struct{}

func (s stdio) Read(b []byte) (int, error) {
	return os.Stdin.Read(b)
}

func (s stdio) Write(b []byte) (int, error) {
	return os.Stdout.Write(b)
}

func (s stdio) Close() error {
	return nil
}
