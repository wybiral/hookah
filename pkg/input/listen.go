package input

import (
	"io"
	"net"

	"github.com/wybiral/hookah/pkg/chreader"
)

type listenApp struct {
	ln net.Listener
	// Channel of messages
	ch chan []byte
	// ch Reader
	r io.Reader
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
	app.r = chreader.New(app.ch)
	go app.serve()
	return app, nil
}

func (app *listenApp) Read(b []byte) (int, error) {
	return app.r.Read(b)
}

func (app *listenApp) Close() error {
	// Closing ch causes r.Read to return EOF
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
