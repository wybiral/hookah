// This package provides the basic input/output parsers for hookah.
package hookah

import (
	"github.com/wybiral/hookah/pkg/input"
	"github.com/wybiral/hookah/pkg/output"
	"io"
)

// Parse an input option string and return a new ReadCloser.
func NewInput(opts string) (io.ReadCloser, error) {
	return input.New(opts)
}

// Parse an output option string and return a new WriteCloser.
func NewOutput(opts string) (io.WriteCloser, error) {
	return output.New(opts)
}
