// Package main implements the main hookah CLI tool.
package main

import (
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
	// Print CLI usage info
	flag.Usage = func() {
		usage(h)
	}
	var rOpts, wOpts flagslice.FlagSlice
	flag.Var(&rOpts, "i", "input node (readonly)")
	flag.Var(&wOpts, "o", "output node (writeonly)")
	flag.Parse()
	config := &app.Config{
		RWOpts: flag.Args(),
		ROpts:  rOpts,
		WOpts:  wOpts,
	}
	err := run(h, config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run(h *hookah.API, config *app.Config) error {
	a, err := app.New(h, config)
	defer a.Close()
	if err != nil {
		return err
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go func() {
		<-interrupt
		a.Quit <- nil
	}()
	return a.Run()
}

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
	os.Exit(0)
}
