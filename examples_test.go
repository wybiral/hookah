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

// Create file input from absolute path.
func ExampleNewInput_fileAbsolute() {
	r, err = hookah.NewInput("file:///relative/path")
}

// Create file input from relative path.
func ExampleNewInput_fileRelative() {
	r, err = hookah.NewInput("file://relative/path")
}

// Create HTTP client input.
func ExampleNewInput_httpClient() {
	r, err = hookah.NewInput("http://localhost:8080")
}

// Create HTTP listen input.
func ExampleNewInput_httpListen() {
	r, err = hookah.NewInput("http-listen://localhost:8080")
}

// Create TCP client input.
func ExampleNewInput_tcpClient() {
	r, err = hookah.NewInput("tcp://localhost:8080")
}

// Create TCP listen input.
func ExampleNewInput_tcpListen() {
	r, err = hookah.NewInput("tcp-listen://localhost:8080")
}

// Create UDP listen input.
func ExampleNewInput_udpListen() {
	r, err = hookah.NewInput("udp-listen://localhost:8080")
}

// Create UDP multicast listen input.
// Here the network interface is set to eth0 but can be ommitted to use the
// default.
func ExampleNewInput_udpMulticast() {
	r, err = hookah.NewInput("udp-multicast://eth0,localhost:8080")
}

// Create Unix client input.
func ExampleNewInput_unixClient() {
	r, err = hookah.NewInput("unix://path/to/sock")
}

// Create Unix listen input.
func ExampleNewInput_unixListen() {
	r, err = hookah.NewInput("unix-listen://path/to/sock")
}

// Create WebSocket client input.
func ExampleNewInput_wsClient() {
	r, err = hookah.NewInput("ws://localhost:8080")
}

// Create WebSocket listen input.
func ExampleNewInput_wsListen() {
	r, err = hookah.NewInput("ws-listen://localhost:8080")
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

// Create file output from absolute path (opens in append mode).
func ExampleNewOutput_fileAbsolute() {
	w, err = hookah.NewOutput("file:///relative/path")
}

// Create file output from relative path (opens in append mode).
func ExampleNewOutput_fileRelative() {
	w, err = hookah.NewOutput("file://relative/path")
}

// Create HTTP client output.
func ExampleNewOutput_httpClient() {
	w, err = hookah.NewOutput("http://localhost:8080")
}

// Create HTTP listen output.
func ExampleNewOutput_httpListen() {
	w, err = hookah.NewOutput("http-listen://localhost:8080")
}

// Create TCP client output.
func ExampleNewOutput_tcpClient() {
	w, err = hookah.NewOutput("tcp://localhost:8080")
}

// Create TCP listen output.
func ExampleNewOutput_tcpListen() {
	w, err = hookah.NewOutput("tcp-listen://localhost:8080")
}

// Create UDP client output.
func ExampleNewOutput_udpClient() {
	w, err = hookah.NewOutput("udp://localhost:8080")
}

// Create Unix client output.
func ExampleNewOutput_unixClient() {
	w, err = hookah.NewOutput("unix://path/to/sock")
}

// Create Unix listen output.
func ExampleNewOutput_unixListen() {
	w, err = hookah.NewOutput("unix-listen://path/to/sock")
}

// Create WebSocket client output.
func ExampleNewOutput_wsClient() {
	w, err = hookah.NewOutput("ws://localhost:8080")
}

// Create WebSocket listen output.
func ExampleNewOutput_wsListen() {
	w, err = hookah.NewOutput("ws-listen://localhost:8080")
}
