package app

import (
	"io"

	"github.com/wybiral/hookah"
	"github.com/wybiral/hookah/pkg/node"
)

// App manages the main hookah CLI app
type App struct {
	// Hookah API instance
	H *hookah.API
	// Channel for outgoing updates
	Ch chan *Update
	// Channel for quit signal
	Quit chan error
	// All nodes (readers and writers combined)
	Nodes []node.Node
	// Reader nodes (inputs)
	Readers []node.Node
	// Writer nodes (outputs)
	Writers []node.Node
	// Current configuration
	Config *Config
}

// New creates a new App instance from the config.
// Callers should close the App even if there's an error here to make sure any
// nodes that did get created are properly closed.
func New(h *hookah.API, config *Config) (*App, error) {
	if config.BufferSize == 0 {
		config.BufferSize = DefaultBufferSize
	}
	a := &App{
		H:       h,
		Ch:      make(chan *Update),
		Quit:    make(chan error, 1),
		Nodes:   make([]node.Node, 0),
		Readers: make([]node.Node, 0),
		Writers: make([]node.Node, 0),
		Config:  config,
	}
	err := a.createNodes()
	if err != nil {
		// a is returned even if there's an error because some of the nodes may
		// have been created and may need to be closed.
		return a, err
	}
	return a, nil
}

// Run the App and block until complete or Quit channel received.
func (a *App) Run() (err error) {
	// Closing this channel will stop readerloops and writerloop
	done := make(chan struct{})
	go func() {
		err = <-a.Quit
		close(done)
	}()
	for _, n := range a.Readers {
		go a.readerloop(n, done)
	}
	er := a.writerloop(done)
	if er != nil {
		err = er
	}
	return err
}

// Close the App and all nodes that were created.
func (a *App) Close() {
	for _, n := range a.Nodes {
		n.Close()
	}
	close(a.Quit)
}

// Create nodes from config
func (a *App) createNodes() error {
	err := a.createRW(a.Config.RWOpts)
	if err != nil {
		return err
	}
	err = a.createR(a.Config.ROpts)
	if err != nil {
		return err
	}
	err = a.createW(a.Config.WOpts)
	if err != nil {
		return err
	}
	if len(a.Nodes) == 1 {
		err = a.createRW([]string{"stdio"})
		if err != nil {
			return err
		}
	}
	return nil
}

// Create bidirectional nodes
func (a *App) createRW(opts []string) error {
	for _, opt := range opts {
		n, err := a.H.NewNode(opt)
		if err != nil {
			return err
		}
		a.Nodes = append(a.Nodes, n)
		a.Readers = append(a.Readers, n)
		a.Writers = append(a.Writers, n)
	}
	return nil
}

// Create reader (input) nodes
func (a *App) createR(opts []string) error {
	for _, opt := range opts {
		n, err := a.H.NewNode(opt)
		if err != nil {
			return err
		}
		a.Nodes = append(a.Nodes, n)
		a.Readers = append(a.Readers, n)
	}
	return nil
}

// Create writer (output) nodes
func (a *App) createW(opts []string) error {
	for _, opt := range opts {
		n, err := a.H.NewNode(opt)
		if err != nil {
			return err
		}
		a.Nodes = append(a.Nodes, n)
		a.Writers = append(a.Writers, n)
	}
	return nil
}

// Called in a goroutine for each reader node. Continuously calls Read on the
// Node and sends the []byte (wrapped in an Update) to App.Ch. Failure here
// should signal on App.Quit.
func (a *App) readerloop(n node.Node, done chan struct{}) {
	for {
		b := make([]byte, a.Config.BufferSize)
		i, err := n.Read(b)
		if i > 0 {
			select {
			case a.Ch <- &Update{Bytes: b[:i], Source: n}:
			case <-done:
				return
			}
		}
		// EOF shouldn't really be an error but it should exit
		if err == io.EOF {
			a.Quit <- nil
			return
		}
		// Other errors should be reported
		if err != nil {
			a.Quit <- err
			return
		}
	}
}

// Receive updates from App.Ch and Write out to all App.Writers. Doesn't Write
// to the Node if it's the source of the Update (to avoid loops).
func (a *App) writerloop(done chan struct{}) error {
	var u *Update
	for {
		select {
		case u = <-a.Ch:
		case <-done:
			return nil
		}
		for _, n := range a.Writers {
			// Don't send to self
			if u.Source == n {
				continue
			}
			_, err := n.Write(u.Bytes)
			if err != nil {
				return err
			}
		}
	}
}
