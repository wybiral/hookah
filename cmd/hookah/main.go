// This package is the main CLI hookah tool.
package main

import (
	"flag"
	"io"
	"log"
	"os"
	"os/signal"
	"github.com/wybiral/hookah/pkg/input"
	"github.com/wybiral/hookah/pkg/output"
)

func main() {
	var inOpts string
	flag.StringVar(&inOpts, "i", "stdin", "Stream input")
	var outOpts string
	flag.StringVar(&outOpts, "o", "stdout", "Stream output")
	flag.Parse()
	in, err := input.Parse(inOpts)
	if err != nil {
		log.Fatal(err)
	}
	out, err := output.Parse(outOpts)
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	go func(){
		<-ch
		in.Close()
		out.Close()
		os.Exit(1)
	}()
	io.Copy(out, in)
}
