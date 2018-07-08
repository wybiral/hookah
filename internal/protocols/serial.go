package protocols

import (
	"net/url"
	"strconv"

	"github.com/tarm/serial"
	"github.com/wybiral/hookah/pkg/node"
)

// Serial creates a serial output and returns Node
func Serial(device string, opts url.Values) (node.Node, error) {
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
	return s, nil
}
