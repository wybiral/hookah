package output

import (
	"io"
	"net/url"
)

// TCPListen creates a TCP server and returns ReadCloser
func TCPListen(addr string, opts url.Values) (io.WriteCloser, error) {
	return listen("tcp", addr)
}
