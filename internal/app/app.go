// Package app implements main hookah application Node manager.
package app

import (
	"io"

	"github.com/wybiral/hookah/pkg/node"
)

// App manages the main hookah CLI app
type App struct {
	// Channel for outgoing updates
	Ch chan *Update
	// Channel for quit signal
	Quit chan error
	// All nodes
	Nodes []*node.Node
	// Reader nodes (inputs)
	Readers []*node.Node
	// Writer nodes (outputs)
	Writers []*node.Node
	// Current configuration
	Config *Config
}

// New creates a new App instance from the config.
func New(config *Config) *App {
	if config == nil {
		config = &Config{}
	}
	if config.BufferSize == 0 {
		config.BufferSize = DefaultBufferSize
	}
	a := &App{
		Ch:      make(chan *Update),
		Quit:    make(chan error, 1),
		Nodes:   make([]*node.Node, 0),
		Readers: make([]*node.Node, 0),
		Writers: make([]*node.Node, 0),
		Config:  config,
	}
	return a
}

// AddNode appends the node to internal lists for processing in Run.
func (a *App) AddNode(n *node.Node) {
	a.Nodes = append(a.Nodes, n)
	if n.R != nil {
		a.Readers = append(a.Readers, n)
	}
	if n.W != nil {
		a.Writers = append(a.Writers, n)
	}
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
		if n.C != nil {
			n.C.Close()
		}
	}
}

// Called in a goroutine for each reader node. Continuously calls Read on the
// Node and sends the []byte (wrapped in an Update) to App.Ch. Failure here
// should signal on App.Quit.
func (a *App) readerloop(n *node.Node, done chan struct{}) {
	for {
		b := make([]byte, a.Config.BufferSize)
		i, err := n.R.Read(b)
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
			_, err := n.W.Write(u.Bytes)
			if err != nil {
				return err
			}
		}
	}
}
