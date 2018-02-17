package cmd_recv

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
)

func usage() {
	fmt.Println("NAME:")
	fmt.Println("   hookah recv - pipe from node to local stdout\n")
	fmt.Println("USAGE:")
	fmt.Println("   hookah recv addr[:port]")
	fmt.Println("")
}

func Main(args []string) {
	if len(args) != 1 {
		usage()
		return
	}
	addr := args[0]
	url := fmt.Sprintf("ws://%s/out", addr)
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	newline := []byte{'\n'}
	out := bufio.NewWriter(os.Stdout)
	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		out.Write(data)
		out.Write(newline)
		out.Flush()
	}
}
