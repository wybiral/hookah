// Package hookah provides input/output stream constructors.
package hookah

import (
	"errors"
	"io"
	"net/url"
	"strings"
	"sync"

	"github.com/wybiral/hookah/pkg/input"
	"github.com/wybiral/hookah/pkg/output"
)

// API is an instance of the Hookah API.
type API struct {
	mu             sync.Mutex
	inputHandlers  map[string]registeredInput
	outputHandlers map[string]registeredOutput
}

// Required info for a registered input
type registeredInput struct {
	handler input.Handler
	proto   string
	usage   string
}

// Required info for a registered ouput
type registeredOutput struct {
	handler output.Handler
	proto   string
	usage   string
}

// New returns a Hookah API instance with default handlers.
func New() *API {
	api := &API{
		inputHandlers:  make(map[string]registeredInput),
		outputHandlers: make(map[string]registeredOutput),
	}
	api.registerInputs()
	api.registerOutputs()
	return api
}

// NewInput parses an input option string and returns a new ReadCloser.
func (a *API) NewInput(op string) (io.ReadCloser, error) {
	proto, arg, opts, err := parseOptions(op)
	if err != nil {
		return nil, err
	}
	reg, ok := a.inputHandlers[proto]
	if !ok {
		return nil, errors.New("unknown input protocol: " + proto)
	}
	return reg.handler(arg, opts)
}

// NewOutput parses an output option string and returns a new WriteCloser.
func (a *API) NewOutput(op string) (io.WriteCloser, error) {
	proto, arg, opts, err := parseOptions(op)
	if err != nil {
		return nil, err
	}
	reg, ok := a.outputHandlers[proto]
	if !ok {
		return nil, errors.New("unknown output protocol: " + proto)
	}
	return reg.handler(arg, opts)
}

// RegisterInput registers a new input protocol.
func (a *API) RegisterInput(proto, usage string, h input.Handler) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.inputHandlers[proto] = registeredInput{
		handler: h,
		proto:   proto,
		usage:   usage,
	}
}

// RegisterOutput registers a new output protocol.
func (a *API) RegisterOutput(proto, usage string, h output.Handler) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.outputHandlers[proto] = registeredOutput{
		handler: h,
		proto:   proto,
		usage:   usage,
	}
}

func (a *API) registerInputs() {
	a.RegisterInput("file", "file://path/to/file", input.File)
	a.RegisterInput("http", "http://address", input.HTTP)
	a.RegisterInput("https", "https://address", input.HTTPS)
	a.RegisterInput("http-listen", "http-listen://address", input.HTTPListen)
	a.RegisterInput("serial", "serial://device?baud=baudrate", input.Serial)
	a.RegisterInput("stdin", "stdin", input.Stdin)
	a.RegisterInput("tcp", "tcp://address", input.TCP)
	a.RegisterInput("tcp-listen", "tcp-listen://address", input.TCPListen)
	a.RegisterInput("udp-listen", "udp-listen://address", input.UDPListen)
	a.RegisterInput("udp-multicast", "udp-multicast://iface,address", input.UDPMulticast)
	a.RegisterInput("unix", "unix://path/to/sock", input.Unix)
	a.RegisterInput("unix-listen", "unix-listen://path/to/sock", input.UnixListen)
	a.RegisterInput("ws", "ws://address", input.WS)
	a.RegisterInput("wss", "wss://address", input.WSS)
	a.RegisterInput("ws-listen", "ws-listen://address", input.WSListen)
}

func (a *API) registerOutputs() {
	a.RegisterOutput("file", "file://path/to/file", output.File)
	a.RegisterOutput("http", "http://address", output.HTTP)
	a.RegisterOutput("https", "https://address", output.HTTPS)
	a.RegisterOutput("http-listen", "http-listen://address", output.HTTPListen)
	a.RegisterOutput("stderr", "stderr", output.Stderr)
	a.RegisterOutput("stdout", "stdout", output.Stdout)
	a.RegisterOutput("tcp", "tcp://address", output.TCP)
	a.RegisterOutput("tcp-listen", "tcp-listen://address", output.TCPListen)
	a.RegisterOutput("udp", "udp://address", output.UDP)
	a.RegisterOutput("unix", "unix://path/to/sock", output.Unix)
	a.RegisterOutput("unix-listen", "unix-listen://path/to/sock", output.UnixListen)
	a.RegisterOutput("ws", "ws://address", output.WS)
	a.RegisterOutput("wss", "wss://address", output.WSS)
	a.RegisterOutput("ws-listen", "ws-listen://address", output.WSListen)
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
