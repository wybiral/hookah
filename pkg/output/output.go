// Package output provides output stream destinations.
package output

import (
	"io"
)

// Handler is the function type for user defined input protocols.
type Handler func(arg string) (io.WriteCloser, error)

// Number of buffered messages for each incoming server connection.
const queueSize = 10
