package protocols

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/wybiral/hookah/pkg/fanout"
	"github.com/wybiral/hookah/pkg/node"
)

type wsListenApp struct {
	sync.Mutex
	server *http.Server
	fan    *fanout.Fanout
	// Channel of messages
	ch chan []byte
	b  []byte
}

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WSListen creates a WebSocket listener Node
func WSListen(addr string) (*node.Node, error) {
	app := &wsListenApp{}
	app.server = &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(app.handle),
	}
	app.fan = fanout.New()
	app.ch = make(chan []byte)
	go app.server.ListenAndServe()
	return node.New(app), nil
}

// WSSListen creates a wss WebSocket listener Node
func WSSListen(arg string) (*node.Node, error) {
	var opts url.Values
	// Parse options
	addrOpts := strings.SplitN(arg, "?", 2)
	addr := addrOpts[0]
	if len(addrOpts) == 2 {
		op, err := url.ParseQuery(addrOpts[1])
		if err != nil {
			return nil, err
		}
		opts = op
	}
	// Load cert and key files
	cert := opts.Get("cert")
	key := opts.Get("key")
	app := &wsListenApp{}
	app.server = &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(app.handle),
	}
	app.fan = fanout.New()
	app.ch = make(chan []byte)
	go app.server.ListenAndServeTLS(cert, key)
	return node.New(app), nil
}

func (app *wsListenApp) Read(b []byte) (int, error) {
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

func (app *wsListenApp) Write(b []byte) (int, error) {
	app.fan.Send(b)
	return len(b), nil
}

func (app *wsListenApp) Close() error {
	return app.server.Close()
}

func (app *wsListenApp) handle(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go app.reader(ws, wg)
	go app.writer(ws, wg)
	wg.Wait()
}

// Pump WebSocket messages to app.ch
func (app *wsListenApp) reader(ws *websocket.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			return
		}
		if len(msg) > 0 {
			app.ch <- msg
		}
	}
}

// Pump app.fan messages to WebSocket
func (app *wsListenApp) writer(ws *websocket.Conn, wg *sync.WaitGroup) {
	ch := make(chan []byte, queueSize)
	app.fan.Add(ch)
	defer func() {
		app.fan.Remove(ch)
		wg.Done()
	}()
	for chunk := range ch {
		err := ws.WriteMessage(websocket.TextMessage, chunk)
		if err != nil {
			return
		}
	}
}
