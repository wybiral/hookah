package node

import (
	"io"
)

// Node is an alias for ReadWriteCloser
type Node io.ReadWriteCloser
