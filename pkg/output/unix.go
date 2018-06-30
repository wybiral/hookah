package output

import (
	"io"
	"net"
	"net/url"
)

// Unix creates a Unix client and returns WriteCloser
func Unix(path string, opts url.Values) (io.WriteCloser, error) {
	return net.Dial("unix", path)
}
