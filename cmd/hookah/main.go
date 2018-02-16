package main

import (
	"fmt"
	"github.com/wybiral/hookah/internal/cmd_node"
	"github.com/wybiral/hookah/internal/cmd_recv"
	"github.com/wybiral/hookah/internal/cmd_send"
	"os"
)

func usage() {
	fmt.Println("NAME:")
	fmt.Println("   hookah - WebSocket pipeline tool\n")
	fmt.Println("USAGE:")
	fmt.Println("   hookah command [command options]\n")
	fmt.Println("COMMANDS:")
	fmt.Println("   node		start new node")
	fmt.Println("   recv		pipe from node to local stdout")
	fmt.Println("   send		pipe local stdout to node")
	fmt.Println("")
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		usage()
		return
	}
	switch args[0] {
	case "node":
		cmd_node.Main(args[1:])
	case "recv":
		cmd_recv.Main(args[1:])
	case "send":
		cmd_send.Main(args[1:])
	default:
		usage()
		return
	}
}
