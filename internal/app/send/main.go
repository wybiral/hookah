package send

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
	fmt.Println("   hookah send - pipe local stdout to node\n")
	fmt.Println("USAGE:")
	fmt.Println("   hookah send addr[:port]")
	fmt.Println("")
}

func Main(args []string) {
	if len(args) != 1 {
		usage()
		return
	}
	url := args[0]
	if !strings.HasPrefix(url, "ws://") {
		url = "ws://" + url
	}
	header := http.Header{"Hookah-Method": {"send"}}
	ws, _, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := scanner.Bytes()
		err := ws.WriteMessage(websocket.BinaryMessage, data)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Sent", len(data), "bytes")
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}
