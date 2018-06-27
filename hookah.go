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
