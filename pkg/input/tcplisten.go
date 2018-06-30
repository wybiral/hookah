package input

import (
	"io"
	"net/url"
)

// TCPListen creates a TCP listener and returns ReadCloser
func TCPListen(addr string, opts url.Values) (io.ReadCloser, error) {
	return listen("tcp", addr)
}
