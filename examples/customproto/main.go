// Example of using hookah with a custom protocol. In this case the protocol is
// named numbers:// and it can accept "odd" or "even" as the argument.
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/wybiral/hookah"
	"github.com/wybiral/hookah/pkg/node"
)

func main() {
	// Create hookah API instance
	h := hookah.New()
	// Register new protocol
	h.RegisterProtocol("numbers", "numbers://parity", numbersHandler)
	// Create hookah node (using our new numbers:// protocol)
	r, err := h.NewNode("numbers://odd")
	if err != nil {
		log.Fatal(err)
	}
	// Create hookah node (stdout)
	w, err := h.NewNode("stdout")
	if err != nil {
		log.Fatal(err)
	}
	// Copy forever
	io.Copy(w.W, r.R)
}

// type to implement Reader interface on.
type numbers int64

// Handlers take an arg string and return a Node
func numbersHandler(arg string) (*node.Node, error) {
	var counter numbers
	if arg == "odd" {
		counter = 1
	} else if arg == "even" {
		counter = 2
	} else {
		return nil, errors.New("numbers requires: odd or even")
	}
	// Node can have R: Reader, W: Writer, C: Closer
	// In this case it's just a Reader
	return &node.Node{R: &counter}, nil
}

// Read the next number (after delay) and increment counter.
func (num *numbers) Read(b []byte) (int, error) {
	// Artificial delay
	time.Sleep(time.Second)
	// Format counter
	s := fmt.Sprintf("%d\n", *num)
	// Increment counter
	*num += 2
	// Copy to byte array
	n := copy(b, []byte(s))
	return n, nil
}
