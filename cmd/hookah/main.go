// This package is the main CLI hookah tool.
package main

import (
	"flag"
	"io"
	"log"
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
	io.Copy(out, in)
}
