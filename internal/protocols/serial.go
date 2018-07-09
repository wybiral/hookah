package protocols

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/tarm/serial"
	"github.com/wybiral/hookah/pkg/node"
)

// Serial creates a serial node
func Serial(arg string) (*node.Node, error) {
	var opts url.Values
	devopts := strings.SplitN(arg, "?", 2)
	device := devopts[0]
	if len(devopts) == 2 {
		op, err := url.ParseQuery(devopts[1])
		if err != nil {
			return nil, err
		}
		opts = op
	}
	baudstr := opts.Get("baud")
	if len(baudstr) == 0 {
		baudstr = "9600"
	}
	baud, err := strconv.ParseInt(baudstr, 10, 32)
	if err != nil {
		return nil, err
	}
	c := &serial.Config{
		Name: device,
		Baud: int(baud),
	}
	s, err := serial.OpenPort(c)
	if err != nil {
		return nil, err
	}
	return node.New(s), nil
}
