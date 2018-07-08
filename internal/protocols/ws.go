package protocols

import (
	"io"
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/wybiral/hookah/pkg/node"
)

type wsconn struct {
	// WebSocket connection
	conn *websocket.Conn
	// Lock for reader
	rmu sync.Mutex
	// Current active reader
	reader io.Reader
	// Lock for writes
	wmu sync.Mutex
}

// WS creates a WebSocket client and returns Node
func WS(addr string, opts url.Values) (node.Node, error) {
	return wsrequest("ws://" + addr)
}

// WSS creates a secure WebSocket client and returns Node
func WSS(addr string, opts url.Values) (node.Node, error) {
	return wsrequest("wss://" + addr)
}

func wsrequest(addr string) (node.Node, error) {
	conn, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		return nil, err
	}
	ws := &wsconn{
		conn: conn,
	}
	return ws, nil
}

func (ws *wsconn) Read(b []byte) (int, error) {
	ws.rmu.Lock()
	defer ws.rmu.Unlock()
	if ws.reader == nil {
		_, reader, err := ws.conn.NextReader()
		if err != nil {
			return 0, err
		}
		ws.reader = reader
	}
	n, err := ws.reader.Read(b)
	if err == io.EOF {
		ws.reader = nil
	} else if err != nil {
		return n, err
	}
	return n, nil
}

func (ws *wsconn) Write(b []byte) (int, error) {
	w, err := ws.conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return 0, err
	}
	defer w.Close()
	return w.Write(b)
}

func (ws *wsconn) Close() error {
	return ws.conn.Close()
}
