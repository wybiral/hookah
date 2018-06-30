package input

import (
	"io"
	"net/url"
)

// UnixListen creates a Unix domain socket listener and returns ReadCloser
func UnixListen(path string, opts url.Values) (io.ReadCloser, error) {
	return listen("unix", path)
}
