package input

import (
	"io"
	"net/url"
)

// TCPListen creates a TCP listener and returns ReadCloser
func TCPListen(path string, args url.Values) (io.ReadCloser, error) {
	return listen("tcp", path)
}
