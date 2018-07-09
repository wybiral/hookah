package protocols

import (
	"errors"
	"net"
	"sync"

	"github.com/wybiral/hookah/pkg/fanout"
	"github.com/wybiral/hookah/pkg/node"
)

type listenApp struct {
	sync.Mutex
	ln  net.Listener
	fan *fanout.Fanout
	// Channel of messages
	ch chan []byte
	b  []byte
}

// listen creates a generic listener and returns Node
func listen(network, addr string) (*node.Node, error) {
	app := &listenApp{}
	ln, err := net.Listen(network, addr)
	if err != nil {
		return nil, err
	}
	app.ln = ln
	app.fan = fanout.New()
	app.ch = make(chan []byte)
	go app.serve()
	return node.New(app), nil
}

func (app *listenApp) Read(b []byte) (int, error) {
	app.Lock()
	defer app.Unlock()
	if len(app.b) == 0 {
		app.b = <-app.ch
	}
	if len(app.b) == 0 {
		return 0, errors.New("listen channel closed")
	}
	n := copy(b, app.b)
	app.b = app.b[n:]
	return n, nil
}

func (app *listenApp) Write(b []byte) (int, error) {
	app.fan.Send(b)
	return len(b), nil
}

func (app *listenApp) Close() error {
	return app.ln.Close()
}

func (app *listenApp) serve() {
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
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go app.reader(conn, wg)
	go app.writer(conn, wg)
	wg.Wait()
}

func (app *listenApp) reader(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		b := make([]byte, bufferSize)
		n, err := conn.Read(b)
		if n > 0 {
			app.ch <- b[:n]
		}
		if err != nil {
			return
		}
	}
}

func (app *listenApp) writer(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
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
