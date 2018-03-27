package app

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
)

func sendUsage() {
	fmt.Println("NAME:")
	fmt.Println("   hookah send - pipe local stdout to node\n")
	fmt.Println("USAGE:")
	fmt.Println("   hookah send addr[:port]")
	fmt.Println("")
}

func SendMain(args []string) {
	if len(args) != 1 {
		sendUsage()
		return
	}
	ws, err := websocketClient(args[0], "send")
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
