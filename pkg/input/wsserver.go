package input

import (
	"io"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type wsServerApp struct {
	server *http.Server
	// Channel of messages
	ch chan []byte
	// Lock for top
	mu *sync.Mutex
	// Current message being read
	top []byte
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Create an WebSocket server and return as ReadCloser
func wsServer(addr string) (io.ReadCloser, error) {
	app := &wsServerApp{}
	app.server = &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(app.handle),
	}
	app.ch = make(chan []byte)
	app.mu = &sync.Mutex{}
	app.top = nil
	go app.server.ListenAndServe()
	return app, nil
}

func (app *wsServerApp) Read(b []byte) (int, error) {
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

func (app *wsServerApp) Close() error {
	close(app.ch)
	return app.server.Close()
}

func (app *wsServerApp) handle(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer ws.Close()
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
