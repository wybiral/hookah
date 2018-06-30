package output

import (
	"io"
	"net/url"
)

// TCPListen creates a TCP server and returns ReadCloser
func TCPListen(path string, args url.Values) (io.WriteCloser, error) {
	return listen("tcp", path)
}
