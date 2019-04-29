// Package hookah provides input/output stream constructors.
package hookah

import (
	"errors"
	"sort"
	"strings"
	"sync"

	"github.com/wybiral/hookah/internal/protocols"
	"github.com/wybiral/hookah/pkg/node"
)

// Version of hookah API.
const Version = "2.0.0"

// API is an instance of the Hookah API.
type API struct {
	mu        sync.RWMutex
	protocols map[string]Protocol
}

// Handler is a function that returns a new Node.
type Handler func(arg string) (*node.Node, error)

// Protocol represents a registered protocol handler.
type Protocol struct {
	Handler Handler
	Proto   string
	Usage   string
}

// New returns a Hookah API instance with default handlers.
func New() *API {
	api := &API{
		protocols: make(map[string]Protocol),
	}
	api.registerProtocols()
	return api
}

// NewNode parses an option string and returns a new Node.
func (a *API) NewNode(op string) (*node.Node, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	proto, arg := parseOptions(op)
	p, ok := a.protocols[proto]
	if !ok {
		return nil, errors.New("unknown protocol: " + proto)
	}
	return p.Handler(arg)
}

// ListProtocols returns all registered protocols.
func (a *API) ListProtocols() []Protocol {
	a.mu.RLock()
	defer a.mu.RUnlock()
	out := make([]Protocol, 0, len(a.protocols))
	for _, p := range a.protocols {
		out = append(out, p)
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].Proto < out[j].Proto
	})
	return out
}

// RegisterProtocol registers a new protocol handler.
func (a *API) RegisterProtocol(proto, usage string, h Handler) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.protocols[proto] = Protocol{
		Handler: h,
		Proto:   proto,
		Usage:   usage,
	}
}

func (a *API) registerProtocols() {
	a.RegisterProtocol("exec", "exec://command", protocols.Exec)
	a.RegisterProtocol("file", "file://path/to/file", protocols.File)
	a.RegisterProtocol("serial", "serial://device?baud=baudrate", protocols.Serial)
	a.RegisterProtocol("stderr", "stderr", protocols.Stderr)
	a.RegisterProtocol("stdin", "stdin", protocols.Stdin)
	a.RegisterProtocol("stdio", "stdio", protocols.Stdio)
	a.RegisterProtocol("stdout", "stdout", protocols.Stdout)
	a.RegisterProtocol("tcp", "tcp://address", protocols.TCP)
	a.RegisterProtocol("tcp-listen", "tcp-listen://address", protocols.TCPListen)
	a.RegisterProtocol("tls", "tls://address?cert=path&insecure=false", protocols.TLS)
	a.RegisterProtocol("tls-listen", "tls-listen://address?cert=path&key=path", protocols.TLSListen)
	a.RegisterProtocol("unix", "unix://path/to/sock", protocols.Unix)
	a.RegisterProtocol("unix-listen", "unix-listen://path/to/sock", protocols.UnixListen)
	a.RegisterProtocol("ws", "ws://address", protocols.WS)
	a.RegisterProtocol("wss", "wss://address", protocols.WSS)
	a.RegisterProtocol("ws-listen", "ws-listen://address", protocols.WSListen)
}

func parseOptions(op string) (string, string) {
	protoarg := strings.SplitN(op, "://", 2)
	if len(protoarg) == 0 {
		return "", ""
	}
	proto := protoarg[0]
	if len(protoarg) == 1 {
		return proto, ""
	}
	return proto, protoarg[1]
}
