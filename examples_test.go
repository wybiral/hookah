package hookah_test

import (
	"io"

	"github.com/wybiral/hookah"
)

var r io.Reader
var w io.Writer
var err error

// NewInput examples

// Create stdin input.
func ExampleNewInput_stdin() {
	r, err = hookah.NewInput("stdin")
}

// Create HTTP client input.
func ExampleNewInput_httpClient() {
	r, err = hookah.NewInput("http://localhost:8080")
}

// Create HTTP server input.
func ExampleNewInput_httpServer() {
	r, err = hookah.NewInput("http-server://localhost:8080")
}

// Create TCP client input.
func ExampleNewInput_tcpClient() {
	r, err = hookah.NewInput("tcp://localhost:8080")
}

// Create TCP server input.
func ExampleNewInput_tcpServer() {
	r, err = hookah.NewInput("tcp-server://localhost:8080")
}

// Create Unix client input.
func ExampleNewInput_unixClient() {
	r, err = hookah.NewInput("unix://path/to/sock")
}

// Create Unix server input.
func ExampleNewInput_unixServer() {
	r, err = hookah.NewInput("unix-server://path/to/sock")
}

// Create WebSocket client input.
func ExampleNewInput_wsClient() {
	r, err = hookah.NewInput("ws://localhost:8080")
}

// Create WebSocket server input.
func ExampleNewInput_wsServer() {
	r, err = hookah.NewInput("ws-server://localhost:8080")
}

// NewOutput examples

// Create stdout output.
func ExampleNewOutput_stdout() {
	w, err = hookah.NewOutput("stdout")
}

// Create stderr output.
func ExampleNewOutput_stderr() {
	w, err = hookah.NewOutput("stderr")
}

// Create HTTP client output.
func ExampleNewOutput_httpClient() {
	w, err = hookah.NewOutput("http://localhost:8080")
}

// Create HTTP server output.
func ExampleNewOutput_httpServer() {
	w, err = hookah.NewOutput("http-server://localhost:8080")
}

// Create TCP client output.
func ExampleNewOutput_tcpClient() {
	w, err = hookah.NewOutput("tcp://localhost:8080")
}

// Create TCP server output.
func ExampleNewOutput_tcpServer() {
	w, err = hookah.NewOutput("tcp-server://localhost:8080")
}

// Create Unix client output.
func ExampleNewOutput_unixClient() {
	w, err = hookah.NewOutput("unix://path/to/sock")
}

// Create Unix server output.
func ExampleNewOutput_unixServer() {
	w, err = hookah.NewOutput("unix-server://path/to/sock")
}

// Create WebSocket client output.
func ExampleNewOutput_wsClient() {
	w, err = hookah.NewOutput("ws://localhost:8080")
}

// Create WebSocket server output.
func ExampleNewOutput_wsServer() {
	w, err = hookah.NewOutput("ws-server://localhost:8080")
}
