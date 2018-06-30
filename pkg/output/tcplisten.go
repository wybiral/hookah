package output

import (
	"io"
)

// TCPListen creates a TCP server and returns ReadCloser
func TCPListen(addr string) (io.WriteCloser, error) {
	return listen("tcp", addr)
}
