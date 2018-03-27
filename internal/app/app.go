package app

import (
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
)

func websocketClient(url, method string) (*websocket.Conn, error) {
	if !(strings.HasPrefix(url, "ws://") || strings.HasPrefix(url, "wss://")) {
		url = "ws://" + url
	}
	header := http.Header{"Hookah-Method": {method}}
	ws, _, err := websocket.DefaultDialer.Dial(url, header)
	return ws, err
}
