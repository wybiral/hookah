// Package output provides output stream destinations.
package output

import (
	"io"
	"net/url"
)

// Handler is the function type for user defined input protocols.
type Handler func(arg string, opts url.Values) (io.WriteCloser, error)

// Number of buffered messages for each incoming server connection.
const queueSize = 10
