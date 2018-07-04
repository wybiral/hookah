package output

import (
	"io"
	"net/url"
	"strconv"

	"github.com/tarm/serial"
)

// Serial creates a serial output and returns WriteCloser
func Serial(device string, opts url.Values) (io.WriteCloser, error) {
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
	return serial.OpenPort(c)
}
