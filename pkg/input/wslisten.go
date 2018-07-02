package input

import (
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/wybiral/hookah/pkg/chreader"
)

type wsListenApp struct {
	server *http.Server
	// Channel of messages
	ch chan []byte
	// ch Reader
	r io.Reader
}

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WSListen creates a WebSocket listener and returns ReadCloser
func WSListen(addr string, opts url.Values) (io.ReadCloser, error) {
	app := &wsListenApp{}
	app.server = &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(app.handle),
	}
	app.ch = make(chan []byte)
	app.r = chreader.New(app.ch)
	go app.server.ListenAndServe()
	return app, nil
}

func (app *wsListenApp) Read(b []byte) (int, error) {
	return app.r.Read(b)
}

func (app *wsListenApp) Close() error {
	// Closing ch causes r.Read to return EOF
	close(app.ch)
	return app.server.Close()
}

func (app *wsListenApp) handle(w http.ResponseWriter, r *http.Request) {
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
