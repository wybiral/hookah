// Example of using hookah with a custom input protocol. In this case the input
// protocol is named numbers:// and it can accept "odd" or "even" as the
// argument.
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/wybiral/hookah"
)

func main() {
	// Create hookah API instance
	h := hookah.New()
	// Register new protocol
	h.RegisterInput("numbers", "numbers://parity", numbersHandler)
	// Create hookah input (using new numbers:// protocol)
	r, err := h.NewInput("numbers://odd")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	// Create hookah output (stdout)
	w, err := h.NewOutput("stdout")
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()
	// Copy forever
	io.Copy(w, r)
}

// struct type to implement interface on.
type numbers struct {
	counter int64
}

// Input handlers take an arg string and return an io.ReadCloser for the input
// stream (or an error).
func numbersHandler(arg string) (io.ReadCloser, error) {
	var counter int64
	if arg == "odd" {
		counter = 1
	} else if arg == "even" {
		counter = 2
	} else {
		return nil, errors.New("numbers requires: odd or even")
	}
	return &numbers{counter: counter}, nil
}

// Read method satisfies the io.ReadCloser interface
func (num *numbers) Read(b []byte) (int, error) {
	// Artificial delay
	time.Sleep(time.Second)
	// Format counter
	s := fmt.Sprintf("%d\n", num.counter)
	// Increment counter
	num.counter += 2
	// Copy to byte array
	n := copy(b, []byte(s))
	return n, nil
}

// Close method satisfies the io.ReadCloser interface
func (num *numbers) Close() error {
	return nil
}
