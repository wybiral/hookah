package input

import (
	"errors"
	"io"
	"net/url"
	"strconv"

	"github.com/tarm/serial"
)

// Serial creates a serial input and returns ReadCloser
func Serial(device string, opts url.Values) (io.ReadCloser, error) {
	baudstr := opts.Get("baud")
	if len(baudstr) == 0 {
		return nil, errors.New("required: baud")
	}
	baud, err := strconv.ParseInt(baudstr, 10, 32)
	if err != nil {
		return nil, err
	}
	config := &serial.Config{
		Name: device,
		Baud: int(baud),
	}
	return serial.OpenPort(config)
}
