package input

import (
	"io"
	"net/http"
)

// HTTP creates a streaming HTTP client and returns ReadCloser
func HTTP(addr string) (io.ReadCloser, error) {
	return httprequest("http://" + addr)
}

// HTTPS creates a streaming HTTPS client and returns as ReadCloser
func HTTPS(addr string) (io.ReadCloser, error) {
	return httprequest("https://" + addr)
}

func httprequest(addr string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}
