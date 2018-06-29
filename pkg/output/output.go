// Package output provides output stream destinations.
package output

import (
	"errors"
	"io"
	"os"
	"strings"
	"sync"
)

// User defined protocols
var protocols sync.Map

// Handler is the function type for user defined input protocols.
type Handler func(arg string) (io.WriteCloser, error)

// Number of buffered messages for each incoming server connection.
const queueSize = 10

// New parses an option string and returns a new WriteCloser.
func New(opts string) (io.WriteCloser, error) {
	parts := strings.SplitN(opts, "://", 2)
	proto := parts[0]
	arg := ""
	if len(parts) > 1 {
		arg = parts[1]
	}
	fn, ok := protocols.Load(proto)
	if ok {
		return fn.(Handler)(arg)
	}
	switch proto {
	case "stdout":
		return os.Stdout, nil
	case "stderr":
		return os.Stderr, nil
	case "file":
		if arg == "" {
			return nil, errors.New("file: no path supplied")
		}
		return file(arg)
	case "http":
		if arg == "" {
			return nil, errors.New("http client: no address supplied")
		}
		return httpClient("http://" + arg)
	case "https":
		if arg == "" {
			return nil, errors.New("https client: no address supplied")
		}
		return httpClient("https://" + arg)
	case "http-listen", "http-server":
		if arg == "" {
			return nil, errors.New("http server: no address supplied")
		}
		return httpServer(arg)
	case "tcp":
		if arg == "" {
			return nil, errors.New("tcp client: no address supplied")
		}
		return tcpClient(arg)
	case "tcp-listen", "tcp-server":
		if arg == "" {
			return nil, errors.New("tcp server: no address supplied")
		}
		return tcpServer(arg)
	case "unix":
		if arg == "" {
			return nil, errors.New("unix client: no address supplied")
		}
		return unixClient(arg)
	case "unix-listen", "unix-server":
		if arg == "" {
			return nil, errors.New("unix server: no address supplied")
		}
		return unixServer(arg)
	case "ws":
		if arg == "" {
			return nil, errors.New("ws client: no address supplied")
		}
		return wsClient("ws://" + arg)
	case "wss":
		if arg == "" {
			return nil, errors.New("wss client: no address supplied")
		}
		return wsClient("wss://" + arg)
	case "ws-listen", "ws-server":
		if arg == "" {
			return nil, errors.New("ws server: no address supplied")
		}
		return wsServer(arg)
	default:
		return nil, errors.New("unknown out protocol: " + proto)
	}
}

// Register a new output protocol.
func Register(proto string, fn Handler) {
	protocols.Store(proto, fn)
}
