// Package hookah provides input/output stream constructors.
package hookah

import (
	"errors"
	"net/url"
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
type Handler func(arg string, opts url.Values) (node.Node, error)

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
func (a *API) NewNode(op string) (node.Node, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	proto, arg, opts, err := parseOptions(op)
	if err != nil {
		return nil, err
	}
	p, ok := a.protocols[proto]
	if !ok {
		return nil, errors.New("unknown protocol: " + proto)
	}
	return p.Handler(arg, opts)
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
	a.RegisterProtocol("file", "file://path/to/file", protocols.File)
	a.RegisterProtocol("serial", "serial://device?baud=baudrate", protocols.Serial)
	a.RegisterProtocol("stderr", "stderr", protocols.Stderr)
	a.RegisterProtocol("stdin", "stdin", protocols.Stdin)
	a.RegisterProtocol("stdio", "stdio", protocols.Stdio)
	a.RegisterProtocol("stdout", "stdout", protocols.Stdout)
	a.RegisterProtocol("tcp", "tcp://address", protocols.TCP)
	a.RegisterProtocol("tcp-listen", "tcp-listen://address", protocols.TCPListen)
	a.RegisterProtocol("unix", "unix://path/to/sock", protocols.Unix)
	a.RegisterProtocol("unix-listen", "unix-listen://path/to/sock", protocols.UnixListen)
	a.RegisterProtocol("ws", "ws://address", protocols.WS)
	a.RegisterProtocol("wss", "wss://address", protocols.WSS)
	a.RegisterProtocol("ws-listen", "ws-listen://address", protocols.WSListen)
}

func parseOptions(op string) (proto, arg string, opts url.Values, err error) {
	protoarg := strings.SplitN(op, "://", 2)
	proto = protoarg[0]
	if len(protoarg) == 1 {
		return
	}
	argopts := strings.SplitN(protoarg[1], "?", 2)
	arg = argopts[0]
	if len(argopts) == 1 {
		return
	}
	opts, err = url.ParseQuery(argopts[1])
	return
}
