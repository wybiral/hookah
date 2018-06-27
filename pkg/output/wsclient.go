package output

import (
	"io"
	"os"
	"sync"

	"github.com/gorilla/websocket"
)

type wsconn struct {
	// WebSocket connection
	conn *websocket.Conn
	// Lock for writer
	mu *sync.Mutex
}

func wsClient(addr string) (io.WriteCloser, error) {
	conn, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		return nil, err
	}
	ws := &wsconn{
		conn: conn,
		mu:   &sync.Mutex{},
	}
	return ws, nil
}

func (ws *wsconn) Write(b []byte) (int, error) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	// When WebSocket is closed conn will be nil
	if ws.conn == nil {
		return 0, os.ErrClosed
	}
	w, err := ws.conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return 0, err
	}
	defer w.Close()
	n, err := w.Write(b)
	return n, err
}

func (ws *wsconn) Close() error {
	var err error
	ws.mu.Lock()
	defer ws.mu.Unlock()
	if ws.conn != nil {
		err = ws.conn.Close()
		ws.conn = nil
	}
	return err
}
