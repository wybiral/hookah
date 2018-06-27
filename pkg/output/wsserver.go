package output

import (
	"io"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/wybiral/hookah/pkg/fanout"
)

type wsServerApp struct {
	server *http.Server
	fan    *fanout.Fanout
}

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Create a WebSocket server and return as ReadCloser
func wsServer(addr string) (io.WriteCloser, error) {
	app := &wsServerApp{}
	app.server = &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(app.handle),
	}
	app.fan = fanout.New()
	go app.server.ListenAndServe()
	return app, nil
}

func (app *wsServerApp) Write(b []byte) (int, error) {
	app.fan.Send(b)
	return len(b), nil
}

func (app *wsServerApp) Close() error {
	return app.server.Close()
}

func (app *wsServerApp) handle(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	go app.writeLoop(ws)
	// Read from connection to process WebSocket control messages
	for {
		_, _, err := ws.NextReader()
		if err != nil {
			return
		}
	}
}

// Register with fanout instance and pump messages to WebSocket client
func (app *wsServerApp) writeLoop(ws *websocket.Conn) {
	ch := make(chan []byte, queueSize)
	app.fan.Add(ch)
	defer func() {
		app.fan.Remove(ch)
		ws.Close()
	}()
	for chunk := range ch {
		err := ws.WriteMessage(websocket.TextMessage, chunk)
		if err != nil {
			return
		}
	}
}
