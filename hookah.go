// Package hookah provides input/output stream constructors.
package hookah

import (
	"io"

	"github.com/wybiral/hookah/pkg/input"
	"github.com/wybiral/hookah/pkg/output"
)

// NewInput parses an input option string and returns a new ReadCloser.
func NewInput(opts string) (io.ReadCloser, error) {
	return input.New(opts)
}

// NewOutput parses an output option string and returns a new WriteCloser.
func NewOutput(opts string) (io.WriteCloser, error) {
	return output.New(opts)
}

// RegisterInput registers a new input protocol.
func RegisterInput(proto string, h input.Handler) {
	input.Register(proto, h)
}

// RegisterOutput registers a new output protocol.
func RegisterOutput(proto string, h output.Handler) {
	output.Register(proto, h)
}
