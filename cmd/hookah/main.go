// Package main implements the main hookah CLI tool.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"

	"github.com/wybiral/hookah"
)

func main() {
	// Create hookah API instance
	h := hookah.New()
	flag.Usage = func() {
		fmt.Print("NAME:\n")
		fmt.Print("   hookah\n\n")
		fmt.Print("USAGE:\n")
		fmt.Print("   hookah -i input -o output\n\n")
		fmt.Print("VERSION:\n")
		fmt.Printf("   %s\n\n", hookah.Version)
		fmt.Print("INPUTS:\n")
		for _, reg := range h.ListInputs() {
			fmt.Printf("   %s\n", reg.Usage)
		}
		fmt.Print("\n")
		fmt.Print("OUTPUTS:\n")
		for _, reg := range h.ListOutputs() {
			fmt.Printf("   %s\n", reg.Usage)
		}
		fmt.Print("\n")
		os.Exit(0)
	}
	// Parse flags
	var inOpts string
	flag.StringVar(&inOpts, "i", "stdin", "Stream input")
	var outOpts string
	flag.StringVar(&outOpts, "o", "stdout", "Stream output")
	flag.Parse()
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
