// This package is the main CLI hookah tool.
package main

import (
	"flag"
	"github.com/wybiral/hookah"
	"io"
	"log"
	"os"
	"os/signal"
)

func main() {
	// Parse flags
	var inOpts string
	flag.StringVar(&inOpts, "i", "stdin", "Stream input")
	var outOpts string
	flag.StringVar(&outOpts, "o", "stdout", "Stream output")
	flag.Parse()
	// Setup in stream
	in, err := hookah.ParseInput(inOpts)
	if err != nil {
		log.Fatal(err)
	}
	// Setup out stream
	out, err := hookah.ParseOutput(outOpts)
	if err != nil {
		log.Fatal(err)
	}
	// Listen for interrupt to close gracefully
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	go func() {
		<-ch
		in.Close()
		out.Close()
		os.Exit(1)
	}()
	// Copy all of in to out
	io.Copy(out, in)
}
