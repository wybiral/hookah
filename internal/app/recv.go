package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func recvUsage() {
	fmt.Println("NAME:")
	fmt.Println("   hookah recv - pipe from node to local stdout\n")
	fmt.Println("USAGE:")
	fmt.Println("   hookah recv addr[:port]")
	fmt.Println("")
}

func RecvMain(args []string) {
	if len(args) != 1 {
		recvUsage()
		return
	}
	ws, err := websocketClient(args[0], "recv")
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
