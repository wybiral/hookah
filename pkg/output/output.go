// This package provides output stream destinations.
package output

import (
	"errors"
	"io"
	"os"
	"strings"
)

// Number of buffered messages for each incoming server connection.
const queueSize = 10

// Parse an option string and return a new WriteCloser.
func New(opts string) (io.WriteCloser, error) {
	parts := strings.SplitN(opts, "://", 2)
	proto := parts[0]
	switch proto {
	case "stdout":
		return os.Stdout, nil
	case "stderr":
		return os.Stderr, nil
	case "http":
		if len(parts) < 2 {
			return nil, errors.New("http client: no address supplied")
		}
		return httpClient(parts[1])
	case "http-server":
		if len(parts) < 2 {
			return nil, errors.New("http server: no address supplied")
		}
		return httpServer(parts[1])
	case "tcp":
		if len(parts) < 2 {
			return nil, errors.New("tcp client: no address supplied")
		}
		return tcpClient(parts[1])
	case "tcp-server":
		if len(parts) < 2 {
			return nil, errors.New("tcp server: no address supplied")
		}
		return tcpServer(parts[1])
	case "unix":
		if len(parts) < 2 {
			return nil, errors.New("unix client: no address supplied")
		}
		return unixClient(parts[1])
	case "unix-server":
		if len(parts) < 2 {
			return nil, errors.New("unix server: no address supplied")
		}
		return unixServer(parts[1])
	default:
		return nil, errors.New("unknown out protocol: " + proto)
	}
}
