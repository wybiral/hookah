package input

import (
	"io"
	"net"
)

// Unix creates a Unix domain socket client and return as ReadCloser
func Unix(addr string) (io.ReadCloser, error) {
	return net.Dial("unix", addr)
}
