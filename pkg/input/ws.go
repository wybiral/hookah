package input

import (
	"io"
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
)

type wsconn struct {
	// WebSocket connection
	conn *websocket.Conn
	// Lock for reader
	mu *sync.Mutex
	// Current active reader
	reader io.Reader
}

// WS Creates a WebSocket client and returns ReadCloser
func WS(addr string, opts url.Values) (io.ReadCloser, error) {
	return wsrequest("ws://" + addr)
}

// WSS Creates a secure WebSocket client and returns ReadCloser
func WSS(addr string, opts url.Values) (io.ReadCloser, error) {
	return wsrequest("wss://" + addr)
}

func wsrequest(addr string) (io.ReadCloser, error) {
	conn, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		return nil, err
	}
	ws := &wsconn{
		conn:   conn,
		mu:     &sync.Mutex{},
		reader: nil,
	}
	return ws, nil
}

func (ws *wsconn) Read(b []byte) (int, error) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	// When WebSocket is closed conn will be nil
	if ws.conn == nil {
		return 0, io.EOF
	}
	// No active reader, get the next one
	if ws.reader == nil {
		_, reader, err := ws.conn.NextReader()
		if err != nil {
			ws.conn.Close()
			ws.conn = nil
			return 0, io.EOF
		}
		ws.reader = reader
	}
	n, err := ws.reader.Read(b)
	// EOF is for this active reader, not socket
	if err == io.EOF {
		ws.reader = nil
	}
	return n, nil
}

func (ws *wsconn) Close() error {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	if ws.conn != nil {
		ws.conn.Close()
		ws.conn = nil
	}
	return nil
}
