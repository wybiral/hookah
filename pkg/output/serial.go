package output

import (
	"errors"
	"io"
	"net/url"
	"strconv"

	"github.com/jacobsa/go-serial/serial"
)

// Serial creates a serial output and returns WriteCloser
func Serial(device string, opts url.Values) (io.WriteCloser, error) {
	baudstr := opts.Get("baud")
	if len(baudstr) == 0 {
		return nil, errors.New("required: baud")
	}
	baud, err := strconv.ParseUint(baudstr, 10, 32)
	if err != nil {
		return nil, err
	}
	options := serial.OpenOptions{
		PortName: device,
		BaudRate: uint(baud),
	}
	return serial.Open(options)
}
