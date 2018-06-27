// Package main implements the main hookah CLI tool.
package main

import (
	"flag"
	"io"
	"log"
	"os"
	"os/signal"

	"github.com/wybiral/hookah"
)

func main() {
	// Parse flags
	var inOpts string
	flag.StringVar(&inOpts, "i", "stdin", "Stream input")
	var outOpts string
	flag.StringVar(&outOpts, "o", "stdout", "Stream output")
	flag.Parse()
	// Setup in stream
	in, err := hookah.NewInput(inOpts)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()
	// Setup out stream
	out, err := hookah.NewOutput(outOpts)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
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
