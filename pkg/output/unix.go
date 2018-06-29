package output

import (
	"io"
	"net"
)

// Unix creates a Unix client and return as WriteCloser
func Unix(addr string) (io.WriteCloser, error) {
	return net.Dial("unix", addr)
}
