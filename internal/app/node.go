package app

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/wybiral/hookah/pkg/fanout"
	"log"
	"net/http"
	"time"
)

// Application state
type node struct {
	fan *fanout.Fanout
}

const writeTimeout = 10 * time.Second // Timeout for client writes
const queueSize = 20                  // Queue size per client
const throttleRate = time.Second / 25 // Maximum rate for incoming messages

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func nodeUsage() {
	fmt.Println("NAME:")
	fmt.Println("   hookah node - start new node\n")
	fmt.Println("USAGE:")
	fmt.Println("   hookah node [addr]:port")
	fmt.Println("")
}

func NodeMain(args []string) {
	if len(args) != 1 {
		nodeUsage()
		return
	}
	n := &node{fan: fanout.New()}
	http.HandleFunc("/", n.handleRequest)
	addr := args[0]
	log.Println("Serving at", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (n *node) handleRequest(w http.ResponseWriter, r *http.Request) {
	method := r.Header.Get("Hookah-Method")
	if method == "send" {
		n.handleSend(w, r)
	} else {
		n.handleRecv(w, r)
	}
}

// Handler for send connections
func (n *node) handleSend(w http.ResponseWriter, r *http.Request) {
	log.Println("Connection: send")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	throttle := time.NewTicker(throttleRate)
	defer throttle.Stop()
	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			return
		}
		log.Println("Received", len(data), "bytes")
		n.fan.Send(data)
		<-throttle.C
	}
}

// Handler for recv connections
func (n *node) handleRecv(w http.ResponseWriter, r *http.Request) {
	log.Println("Connection: recv")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	go n.recvReadLoop(ws)
	n.recvWriteLoop(ws)
}

// Read from recv connections to process WebSocket control messages
func (n *node) recvReadLoop(ws *websocket.Conn) {
	defer ws.Close()
	for {
		_, _, err := ws.NextReader()
		if err != nil {
			return
		}
	}
}

// Register with fanout instance and pump messages to WebSocket client
func (n *node) recvWriteLoop(ws *websocket.Conn) {
	ch := make(chan []byte, queueSize)
	n.fan.Add(ch)
	defer func() {
		n.fan.Remove(ch)
		ws.Close()
	}()
	for msg := range ch {
		ws.SetWriteDeadline(time.Now().Add(writeTimeout))
		err := ws.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
