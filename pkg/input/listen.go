package input

import (
	"io"
	"net"
	"sync"
)

type listenApp struct {
	ln net.Listener
	// Channel of messages
	ch chan []byte
	// Lock for changing top
	mu *sync.Mutex
	// Current message being read
	top []byte
}

// listen creates a generic listener and returns ReadCloser
func listen(network, addr string) (io.ReadCloser, error) {
	app := &listenApp{}
	ln, err := net.Listen(network, addr)
	if err != nil {
		return nil, err
	}
	app.ln = ln
	app.ch = make(chan []byte)
	app.mu = &sync.Mutex{}
	app.top = nil
	go app.serve()
	return app, nil
}

func (app *listenApp) Read(b []byte) (int, error) {
	app.mu.Lock()
	defer app.mu.Unlock()
	if len(app.top) == 0 {
		app.top = <-app.ch
		if len(app.top) == 0 {
			// ch is closed
			return 0, io.EOF
		}
	}
	n := copy(b, app.top)
	app.top = app.top[n:]
	return n, nil
}

func (app *listenApp) Close() error {
	close(app.ch)
	return app.ln.Close()
}

func (app *listenApp) serve() {
	defer app.ln.Close()
	for {
		conn, err := app.ln.Accept()
		if err != nil {
			return
		}
		go app.handle(conn)
	}
}

func (app *listenApp) handle(conn net.Conn) {
	defer conn.Close()
	for {
		b := make([]byte, bufferSize)
		n, err := conn.Read(b)
		if err != nil {
			return
		}
		app.ch <- b[:n]
	}
}
