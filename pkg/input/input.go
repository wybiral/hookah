// Package input provides input stream sources.
package input

import (
	"io"
)

// Handler is the function type for user defined input protocols.
type Handler func(arg string) (io.ReadCloser, error)

// Buffer size used for incoming messages to servers
const bufferSize = 4 * 1024
