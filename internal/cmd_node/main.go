package cmd_node

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/wybiral/hookah/pkg/fanout"
	"log"
	"net/http"
	"time"
)

func usage() {
	fmt.Println("NAME:")
	fmt.Println("   hookah node - start new node\n")
	fmt.Println("USAGE:")
	fmt.Println("   hookah node [addr]:port")
	fmt.Println("")
}

func Main(args []string) {
	if len(args) != 1 {
		usage()
		return
	}
	node := &Node{fan: fanout.NewFanout()}
	r := http.NewServeMux()
	r.HandleFunc("/in", node.In)
	r.HandleFunc("/out", node.Out)
	addr := args[0]
	log.Println("Serving at", addr)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatal(err)
	}
}

type Node struct {
	fan *fanout.Fanout
}

const writeTimeout = 10 * time.Second // Timeout for client writes
const queueSize    = 20
const throttleRate = time.Second / 25

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (node *Node) In(w http.ResponseWriter, r *http.Request) {
	log.Println("/in")
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
		node.fan.Send(data)
		<-throttle.C
	}
}

func (node *Node) Out(w http.ResponseWriter, r *http.Request) {
	log.Println("/out")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	go outReadLoop(ws)
	outWriteLoop(ws, node)
}

func outReadLoop(ws *websocket.Conn) {
	defer ws.Close()
	for {
		_, _, err := ws.NextReader()
		if err != nil {
			return
		}
	}
}

func outWriteLoop(ws *websocket.Conn, node *Node) {
	ch := make(chan []byte, queueSize)
	node.fan.AddChan(ch)
	defer func() {
		node.fan.RemoveChan(ch)
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
