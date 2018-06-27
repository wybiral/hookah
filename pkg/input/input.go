// Package input provides input stream sources.
package input

import (
	"errors"
	"io"
	"os"
	"strings"
)

// Buffer size used for incoming messages to servers
const bufferSize = 4 * 1024

// New parses an option string and returns a new ReadCloser.
func New(opts string) (io.ReadCloser, error) {
	parts := strings.SplitN(opts, "://", 2)
	proto := parts[0]
	switch proto {
	case "stdin":
		return os.Stdin, nil
	case "file":
		if len(parts) < 2 {
			return nil, errors.New("file: no path supplied")
		}
		return file(parts[1])
	case "http":
		if len(parts) < 2 {
			return nil, errors.New("http client: no address supplied")
		}
		return httpClient("http://" + parts[1])
	case "https":
		if len(parts) < 2 {
			return nil, errors.New("https client: no address supplied")
		}
		return httpClient("https://" + parts[1])
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
	case "ws":
		if len(parts) < 2 {
			return nil, errors.New("ws client: no address supplied")
		}
		return wsClient("ws://" + parts[1])
	case "wss":
		if len(parts) < 2 {
			return nil, errors.New("wss client: no address supplied")
		}
		return wsClient("wss://" + parts[1])
	case "ws-server":
		if len(parts) < 2 {
			return nil, errors.New("ws server: no address supplied")
		}
		return wsServer(parts[1])
	default:
		return nil, errors.New("unknown in protocol: " + proto)
	}
}
