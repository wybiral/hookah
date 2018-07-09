// Package main implements the main hookah CLI tool.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/wybiral/hookah"
	"github.com/wybiral/hookah/internal/app"
	"github.com/wybiral/hookah/pkg/flagslice"
)

func main() {
	// Create hookah API instance
	h := hookah.New()
	flag.Usage = func() {
		usage(h)
		os.Exit(0)
	}
	// Parse flags
	var opts, rOpts, wOpts flagslice.FlagSlice
	flag.Var(&rOpts, "i", "input node (readonly)")
	flag.Var(&wOpts, "o", "output node (writeonly)")
	flag.Parse()
	opts = flag.Args()
	// Run and report errors
	err := run(h, opts, rOpts, wOpts)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run(h *hookah.API, opts, rOpts, wOpts []string) error {
	a := app.New(nil)
	// Closing this App instance will close all the nodes being created
	defer a.Close()
	// Add bidirectional nodes
	err := addNodes(h, a, opts, true, true)
	if err != nil {
		return err
	}
	// Add reader (input) nodes
	err = addNodes(h, a, rOpts, true, false)
	if err != nil {
		return err
	}
	// Add writer (output) nodes
	err = addNodes(h, a, wOpts, false, true)
	if err != nil {
		return err
	}
	// No nodes, show usage
	if len(a.Nodes) == 0 {
		flag.Usage()
		return nil
	}
	// Only one node, link to stdio
	if len(a.Nodes) == 1 {
		n, err := h.NewNode("stdio")
		if err != nil {
			return err
		}
		a.AddNode(n)
	}
	if len(a.Readers) == 0 {
		return errors.New("no input nodes")
	}
	if len(a.Writers) == 0 {
		return errors.New("no output nodes")
	}
	// Handle CTRL+C by sending a Quit signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go func() {
		<-interrupt
		a.Quit <- nil
	}()
	return a.Run()
}

// Add nodes to the app from opt strings and r/w status
func addNodes(h *hookah.API, a *app.App, opts []string, r, w bool) error {
	for _, opt := range opts {
		n, err := h.NewNode(opt)
		if err != nil {
			return err
		}
		if !r {
			n.R = nil
		}
		if !w {
			n.W = nil
		}
		a.AddNode(n)
	}
	return nil
}

// Print CLI usage info
func usage(h *hookah.API) {
	fmt.Print("NAME:\n")
	fmt.Print("   hookah\n\n")
	fmt.Print("USAGE:\n")
	fmt.Print("   hookah node [node] -i in_node -o out_node\n\n")
	fmt.Print("VERSION:\n")
	fmt.Printf("   %s\n\n", hookah.Version)
	fmt.Print("PROTOCOLS:\n")
	for _, p := range h.ListProtocols() {
		fmt.Printf("   %s\n", p.Usage)
	}
	fmt.Print("\n")
}
