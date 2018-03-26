package recv

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"strings"
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
	url := args[0]
	if !(strings.HasPrefix(url, "ws://") || strings.HasPrefix(url, "wss://")) {
		url = "ws://" + url
	}
	header := http.Header{"Hookah-Method": {"recv"}}
	ws, _, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	out := bufio.NewWriter(os.Stdout)
	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		out.Write(data)
		out.WriteByte('\n')
		out.Flush()
	}
}
