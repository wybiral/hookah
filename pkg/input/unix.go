package input

import (
	"io"
	"net"
	"net/url"
)

// Unix creates a Unix domain socket client and returns ReadCloser
func Unix(path string, opts url.Values) (io.ReadCloser, error) {
	return net.Dial("unix", path)
}
