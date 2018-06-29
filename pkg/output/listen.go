package output

import (
	"io"
	"net"

	"github.com/wybiral/hookah/pkg/fanout"
)

type listenApp struct {
	ln  net.Listener
	fan *fanout.Fanout
}

// Create a listen server and return as ReadCloser
func listen(network, addr string) (io.WriteCloser, error) {
	app := &listenApp{}
	ln, err := net.Listen(network, addr)
	if err != nil {
		return nil, err
	}
	app.ln = ln
	app.fan = fanout.New()
	go app.serve()
	return app, nil
}

func (app *listenApp) Write(b []byte) (int, error) {
	app.fan.Send(b)
	return len(b), nil
}

func (app *listenApp) Close() error {
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
	ch := make(chan []byte, queueSize)
	app.fan.Add(ch)
	defer app.fan.Remove(ch)
	for chunk := range ch {
		_, err := conn.Write(chunk)
		if err != nil {
			return
		}
	}
}
