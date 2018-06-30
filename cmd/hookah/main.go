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
	// Create hookah API instance
	h := hookah.New()
	// Setup input stream
	r, err := h.NewInput(inOpts)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	// Setup output stream
	w, err := h.NewOutput(outOpts)
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()
	// Listen for interrupt to close gracefully
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	go func() {
		<-ch
		r.Close()
		w.Close()
		os.Exit(1)
	}()
	// Copy all of in to out
	io.Copy(w, r)
}
