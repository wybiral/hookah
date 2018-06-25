package input

import (
	"io"
	"net"
	"sync"
)

type listenServerApp struct {
	ln net.Listener
	// Channel of messages
	ch chan []byte
	// Lock for changing top
	mu *sync.Mutex
	// Current message being read
	top []byte
}

// Create a listen server and return as ReadCloser
func listenServer(network, addr string) (io.ReadCloser, error) {
	app := &listenServerApp{}
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

func (app *listenServerApp) Read(b []byte) (int, error) {
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

func (app *listenServerApp) Close() error {
	close(app.ch)
	return app.ln.Close()
}

func (app *listenServerApp) serve() {
	defer app.ln.Close()
	for {
		conn, err := app.ln.Accept()
		if err != nil {
			return
		}
		go app.handle(conn)
	}
}

func (app *listenServerApp) handle(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, bufferSize)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		app.ch <- buffer[:n]
	}
}
