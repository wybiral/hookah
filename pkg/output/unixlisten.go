package output

import (
	"io"
	"net/url"
)

// UnixListen creates a Unix domain socket listener and returns ReadCloser
func UnixListen(path string, opts url.Values) (io.WriteCloser, error) {
	return listen("unix", path)
}
